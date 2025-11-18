package model

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"senano-music/util"
	"senano-music/util/audiofileutil"
	"senano-music/util/mylog"
	"strconv"
	"sync"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"go.senan.xyz/taglib"
	"gorm.io/gorm"
)

// 基础信息
type MusicInfo struct {
	gorm.Model
	Id         int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	BaseDir    string `json:"basedir"`    // 所在文件夹
	Path       string `json:"path"`       // 绝对路径
	Cover      string `json:"cover"`      // 封面
	Title      string `json:"title"`      // 标题
	Artist     string `json:"artist"`     // 艺术家
	Album      string `json:"album"`      // 专辑
	Comment    string `json:"comment"`    // 简介
	Genre      string `json:"genre"`      // 风格
	Year       int    `json:"year"`       // 年份
	Track      int    `json:"track"`      // 轨道
	Length     int    `json:"length"`     // 时长
	Bitrate    int    `json:"bitrate"`    // 比特率
	Samplerate int    `json:"samplerate"` // 采样率
	Channels   int    `json:"channels"`   // 通道
}

// parse info from music file
func MusicParse(path *string, basedir *string) {
	tags, err := taglib.ReadTags(*path)
	if err != nil {
		mylog.LOG.Error().Msg(fmt.Sprintf("Read tags failed: %s", err.Error()))
		return
	}

	properties, err := taglib.ReadProperties(*path)
	if err != nil {
		mylog.LOG.Error().Msg(fmt.Sprintf("Read properties failed: %s", err.Error()))
		return
	}

	cover := getCover(path)

	title := tags[taglib.Title][0]
	artist := util.FirstOrDefault(tags[taglib.Artist], "")
	album := util.FirstOrDefault(tags[taglib.Album], "")
	comment := util.FirstOrDefault(tags[taglib.Comment], "")
	genre := util.FirstOrDefault(tags[taglib.Genre], "")
	date := util.FirstOrDefault(tags[taglib.Date], "1997-01-01")
	time, _ := time.Parse("yyyy-MM-dd", date)
	year := time.Year()
	track, _ := strconv.Atoi(util.FirstOrDefault(tags[taglib.TrackNumber], "0"))

	musicInfo := MusicInfo{
		Title:      title,
		Path:       *path,
		BaseDir:    *basedir,
		Cover:      cover,
		Artist:     artist,
		Album:      album,
		Comment:    comment,
		Genre:      genre,
		Year:       year,
		Track:      track,
		Length:     int(properties.Length),
		Bitrate:    int(properties.Bitrate),
		Samplerate: int(properties.SampleRate),
		Channels:   int(properties.Channels),
	}

	DB.Create(&musicInfo)
}

// 读取封面
func getCover(musicFilePath *string) string {
	coverFilePath := ""

	// 替换文件后缀，重命名文件名
	fileExt := filepath.Ext(*musicFilePath)
	coverFilePath = audiofileutil.ReplaceLast(*musicFilePath, fileExt, "-cover.jpg")

	// 文件已存在，直接返回
	isExists := audiofileutil.CheckFilePathExist(coverFilePath)
	if isExists {
		return coverFilePath
	}

	// 从文件中读取封面
	coverByteData, err := taglib.ReadImage(*musicFilePath)
	if err != nil {
		mylog.LOG.Warn().Msg(fmt.Sprintf("Read cover failed: %s", err.Error()))
	}

	// 创建文件
	file, err := os.Create(coverFilePath)
	if err != nil {
		mylog.LOG.Warn().Msg(fmt.Sprintf("Error creating file: %s", err.Error()))
	}
	defer file.Close()

	// 写入文件
	_, err = file.Write(coverByteData)
	if err != nil {
		mylog.LOG.Warn().Msg(fmt.Sprintf("Error writing to file: %s", err.Error()))
	}
	return coverFilePath
}

// 清除以前的数据并重新扫描文件夹所有歌曲
func MusicScan(dir *string) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		clearOldMusicInfo()
		doMusicScan(dir)
		return nil
	})
}

// 执行扫描
func doMusicScan(dir *string) error {

	var wg sync.WaitGroup
	// 新建等待队列
	fileChan := make(chan string, 300)

	// 新建工作队列，大小等于cpu数量
	numWorkers := runtime.NumCPU()
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 从队列中取出路径
			for musicPath := range fileChan {
				absPath, err := filepath.Abs(musicPath)
				basedir := filepath.Dir(absPath)
				// 文件路径、文件类型校验
				if err != nil {
					mylog.LOG.Warn().Msg(fmt.Sprintf("获取文件路径失败: %s", err.Error()))
					return
				}
				fileType, err := mimetype.DetectFile(absPath)
				if err != nil {
					mylog.LOG.Debug().Msg(fmt.Sprintf("跳过不支持的文件类型: %s", err.Error()))
					return
				}
				_, err = audiofileutil.GetAudioFileType(fileType.String())
				if err != nil {
					mylog.LOG.Debug().Msg(fmt.Sprintf("跳过不支持的文件类型: %s", fileType.String()))
					return
				}

				// 已经添加过不再扫描
				// musicList := FindMusicByPath(&absPath)
				// if len(musicList) > 0 {
				// 	continue
				// }

				MusicParse(&absPath, &basedir)
			}
		}()
	}

	// 开始遍历目录，放入队列
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// 将扫描到的文件放入队列
		fileChan <- path

		return nil
	})
	if err != nil {
		return err
	}

	close(fileChan)
	wg.Wait()

	// TODO 更新收藏集

	return nil
}

func FindMusicByPath(path *string) []MusicInfo {
	musicList := []MusicInfo{}
	DB.Where("path", path).Find(&musicList)
	return musicList
}

func DeleteMusicByPath(path *string) {
	DB.Unscoped().Where("path", path).Delete(&MusicInfo{})
}

// 清除所有历史数据
func clearOldMusicInfo() {
	DB.Unscoped().Where("1=1").Delete(&MusicInfo{})
}

func MusicList() []MusicInfo {
	musicList := []MusicInfo{}
	DB.Find(&musicList)
	return musicList
}
