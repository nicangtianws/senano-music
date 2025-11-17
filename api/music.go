package api

import (
	"senano-music/model"
)

func (a *App) MusicScan() string {
	musicDir := "/home/swu/projects/wails-demo/senano-music/test/music"
	model.MusicScan(&musicDir)
	return ResultSuccess()
}

func (a *App) MusicList() string {
	musicInfos := model.MusicList()
	return ResultData(musicInfos)
}
