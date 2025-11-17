package api

import (
	"context"
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
	basedir := "/home/swu/projects/wails-demo/senano-music/test"
	model.InitDatabase(&basedir)
	mylog.InitLog(&basedir)

	a.ctx = ctx
}
