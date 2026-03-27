package scraping

import (
	"fmt"
	"path/filepath"
)

func toPNG(name string) string {
	ext := filepath.Ext(name)
	return name[:len(name)-len(ext)] + ".png"
}

func LRGetGameCoverUrl(rom Rom) (url string) {
	fmt.Println(rom.Platform)
	url = "https://thumbnails.libretro.com/" + platformForLR(rom.Platform) + "/Named_Boxarts/" + toPNG(rom.RomName)
	return
}
