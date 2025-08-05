package scraping_test

import (
	"os"
	"slices"
	"testing"

	"github.com/FullStackDev200/RomScraper/scraping"
)

func TestTGDBGetGamesByName(t *testing.T) {
	// Set your API key (better to use os.Setenv in test environment)
	os.Setenv("TGD_API_KEY", "d8a6fe49e00357deaa75b74363f298b9ff9ee954eb4d036cce87f8a6461f23c6")

	tests := []struct {
		name       string
		searchName string
		want       []string
		want2      []int64
	}{
		{
			name:       "Find Super Mario Bros",
			searchName: "Super Mario Bros",
			want:       []string{"Super Mario Bros"},
			want2:      []int64{494}, // Example ID; update with the actual value
		},
		{
			name:       "Find Zelda",
			searchName: "Zelda",
			want:       []string{"The Legend of Zelda"},
			want2:      []int64{295}, // Example ID
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := scraping.TGDBGetGamesByName(tt.searchName)

			if len(got) == 0 {
				t.Errorf("Expected at least one title, got none")
			}
			if len(got2) == 0 {
				t.Errorf("Expected at least one ID, got none")
			}

			found := slices.Contains(got, tt.want[0])
			if !found {
				t.Errorf("Expected title %q not found in results: %v", tt.want[0], got)
			}
		})
	}
}

func TestTGDBGetImageCover(t *testing.T) {

	os.Setenv("TGD_API_KEY", "d8a6fe49e00357deaa75b74363f298b9ff9ee954eb4d036cce87f8a6461f23c6")
	tests := []struct {
		name string
		id   int64
	}{
		{
			name: "Check that Super Mario cover image exists",
			id:   86880,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := scraping.TGDBGetImageCover(tt.id)
			if got == nil || got.Bounds().Empty() {
				t.Errorf("TGDBGetImageCover(%d) returned nil or empty image", tt.id)
				t.Logf("Image dimensions: %dx%d", got.Bounds().Dx(), got.Bounds().Dy())
			}
		})
	}
}
