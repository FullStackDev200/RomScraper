package scraping_test

import (
	"RomScraper/scraping"
	"strings"

	"testing"
)

func TestVimmSearchRoms(t *testing.T) {
	tests := []struct {
		name       string
		searchName string
		filter     string
	}{
		{
			name:       "Find Super Mario Bros Rom's Url",
			searchName: "Super Mario Bros",
			filter:     "GBA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := scraping.VimmSearchRoms(tt.searchName, tt.filter)

			if len(got) == 0 {
				t.Errorf("expected at least one ROM, got none")
			}

			// optional: check if one of the results contains "Super Mario Bros"
			found := false
			for _, rom := range got {
				if strings.Contains(rom.Title, "Super Mario Bros") {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected a ROM with 'Super Mario Bros' in title, got %+v", got)
			}
		})
	}
}

func TestDownloadGame(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		rom  scraping.Rom
		path string
	}{
		{
			name: "Super Mario Download Test",
			rom:  scraping.Rom{DownloadUrl: "https://dl3.vimm.net/?mediaId=4007"},
			path: "~/Downloads/test.7z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraping.DownloadGame(tt.rom, tt.path)
		})
	}
}
