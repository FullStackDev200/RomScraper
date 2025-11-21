package main

import (
	"context"
	"fmt"
	"image"

	sc "github.com/FullStackDev200/RomScraper/scraping"
	uc "github.com/FullStackDev200/RomScraper/userconfig"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetGamesByName(name string) []sc.Game {
	games := sc.TGDBGetGamesByName(name)
	return games
}

func (a *App) GetGameCover(url string) image.Image {
	cover := sc.TGDBGetGameCover(url)
	return cover
}

func (a *App) GetGameCoverUrl(id int64) string {
	cover := sc.TGDBGetGameCoverUrl(id)
	return cover
}

func (a *App) VimSearchGames(gameName string, filter string) []sc.Rom {
	roms := sc.VimmSearchRoms(gameName, filter)
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

func (a *App) NewConfigStore() (*uc.ConfigStore, error) {
	store, err := uc.NewConfigStore()

	if err != nil {
		fmt.Printf("could not initialize the config store: %v\n", err)
		return &uc.ConfigStore{}, err
	}
	return store, err
}

func (a *App) GetConfig(cs *uc.ConfigStore) (uc.Config, error) {
	cfg, err := cs.Get()
	if err != nil {
		fmt.Printf("could not retrieve the configuration: %v\n", err)
		return uc.Config{}, err
	}

	return cfg, err
}

