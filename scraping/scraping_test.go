// scraping/scraping_test.go
package scraping

import (
	"testing"
)

func TestGetPossibleGames(t *testing.T) {
	games := GetPossibleGames("Mario")
	if len(games) == 0 {
		t.Fatal("Expected at least one game")
	}

	for _, game := range games {
		t.Log(game.Title)
		t.Log(game.CoverUrl)
	}
}

func TestDownloadCover(t *testing.T) {
	DownloadCover("https://images.emulatorgames.net/gameboy-advance/super-mario-advance-4-super-mario-bros-3-v1-1.webp", "Mario")
}
