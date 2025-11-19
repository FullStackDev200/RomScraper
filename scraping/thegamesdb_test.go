package scraping_test

import (
	"os"
	"testing"

	"github.com/FullStackDev200/RomScraper/scraping"
)

func TestTGDBGetGamesByName(t *testing.T) {
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
			got := scraping.TGDBGetGamesByName(tt.searchName)

			if len(got) == 0 {
				t.Errorf("Expected at least one title, got none")
			}
		})
	}
}
