package main

import (
	"context"

	sc "github.com/FullStackDev200/RomScraper/scraping"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetGamesByName(name string) []sc.Rom {
	// TODO: Make GetGameCoverUrl part of this
	games := sc.VimmSearchRoms(name)
	return games
}

func (a *App) GetGameCoverUrl(rom sc.Rom) string {
	cover := sc.LRGetGameCoverUrl(rom)
	return cover
}

func (a *App) VimSearchGames(gameName string, filter string) []sc.Rom {
	roms := sc.VimmSearchRoms(gameName)
	return roms
}

func (a *App) VimDownloadGame(rom sc.Rom, path string) {
	sc.DownloadGame(rom, path)
}

func (a *App) ChooseDirectory() (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a directory",
	})
	if err != nil {
		return "", err
	}

	return path, nil
}

func (a *App) RAvalidateHash(rom sc.Rom) bool {
	return sc.RAvalidateHash(rom)
}

func (a *App) PlatformToString(platform sc.Platform) string {
	return sc.PlatformToString(platform)
}
