package api

import (
	"context"
	"os"
	"path/filepath"
	"senano-music/model"
	"senano-music/util/mylog"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	// 生成默认为应用目录
	home, _ := os.UserHomeDir()
	basedir := filepath.Join(home, ".local", "senano-music")
	model.InitDatabase(&basedir)
	mylog.InitLog(&basedir)

	a.ctx = ctx
}
