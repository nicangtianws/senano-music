package audiofileutil

import (
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/strutil"
)

func CheckFilePathExist(checkFilePath string) bool {
	var exist = true
	if _, err := os.Stat(checkFilePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 将基础数据路径转换为绝对路径
// @param path 路径
func AbsBasedir(path string) string {
	if strutil.IsBlank(path) {
		panic("Basedir is blank")
	}

	info, err := os.Stat(path)

	if err != nil && !os.IsExist(err) {
		panic("Basedir not exists: " + path)
	}

	if !info.IsDir() {
		panic("Basedir is not dir: " + path)
	}

	pathAbs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return pathAbs
}

// ReplaceLastN 从后往前替换n次出现
func ReplaceLastN(s, old, new string, n int) string {
	if n <= 0 || old == "" {
		return s
	}

	runes := []rune(s)
	oldRunes := []rune(old)
	newRunes := []rune(new)

	count := 0
	for i := len(runes) - len(oldRunes); i >= 0; i-- {
		if string(runes[i:i+len(oldRunes)]) == old {
			// 替换匹配的部分
			runes = append(runes[:i], append(newRunes, runes[i+len(oldRunes):]...)...)
			count++
			if count >= n {
				break
			}
		}
	}

	return string(runes)
}

// ReplaceLast 替换最后一次出现
func ReplaceLast(s, old, new string) string {
	return ReplaceLastN(s, old, new, 1)
}
