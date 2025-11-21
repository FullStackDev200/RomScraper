package scraping_test

import (
	"RomScraper/scraping"
	"testing"

	"github.com/joho/godotenv"
)

func TestRAgetConsoleGamesList(t *testing.T) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		t.Fatalf("failed to load test env: %v", err)
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rom scraping.Rom
	}{
		{
			name: "Test Valid Hash Rom on Snes",
			rom: scraping.Rom{
				Title:    "Super Mario World",
				Platform: "Super Nintendo",
				// GameHash: "38bb405ba6c6714697b48fb0ad15a2a1",
				GameHash: struct {
					SHA1 string
					MD5  string
					CRC  string
				}{
					SHA1: "38bb405ba6c6714697b48fb0ad15a2a1",
				},
			},
		},

		{
			name: "Test Invalid Hash Rom on Snes",
			rom: scraping.Rom{
				Title:    "Super Mario World",
				Platform: "Super Nintendo",
				GameHash: struct {
					SHA1 string
					MD5  string
					CRC  string
				}{
					SHA1: "38bb405bc6714697b48fb0ad15a2a1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraping.RAvalidateHash(tt.rom)
		})
	}
}
