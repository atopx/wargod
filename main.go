package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"wargod/api"
	"wargod/game"
)

//go:embed all:frontend/dist
var assets embed.FS

// App struct
type App struct {
	ctx context.Context
	api *api.Api
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.api.Ctx = ctx
	a.api.Game = game.New(ctx)
	go func() {
		a.api.Game.Start()
	}()
}

func (a *App) domReady(ctx context.Context) {}

func (a *App) beforeClose(ctx context.Context) bool {
	return false
}

func (a *App) shutdown(ctx context.Context) {}

func (a *App) onSecondInstanceLaunch(data options.SecondInstanceData) {}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func main() {
	// Create an instance of the app structure
	app := &App{api: api.New()}
	// Create application with options
	err := wails.Run(&options.App{
		Title:             "极地战神",
		Width:             900,
		Height:            1200,
		MinWidth:          900,
		MinHeight:         1200,
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer:       &assetserver.Options{Assets: assets},
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		LogLevel:          logger.DEBUG,
		Menu:              nil,
		Logger:            nil,
		WindowStartState:  options.Normal,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "wargod-e3984e08-28dc-4e3d-b70a-45e961589cdc",
			OnSecondInstanceLaunch: app.onSecondInstanceLaunch,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "./",
		},
		Bind: []interface{}{
			app.api,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
