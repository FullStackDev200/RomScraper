package scraping

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/agnivade/levenshtein"
	"github.com/buger/jsonparser"
)

type GameHash struct {
	SHA1 string
	MD5  string
	CRC  string
}

type Rom struct {
	Title       string
	Id          int64
	Platform    Platform
	GameHash    GameHash
	CoverUrl    string
	CoverImg    image.Image
	DownloadUrl string
	PageUrl     string
	RomName     string
}

func VimmSearchRoms(gameName string) (roms []Rom) {
	client := &http.Client{
		Transport: &http.Transport{
			// Disabled verification for vimms
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, err := http.NewRequest("GET", "https://vimm.net/vault/?p=list&q="+url.QueryEscape(gameName), nil)

	if err != nil {
		log.Println(err)
	}

	req.Header.Set("accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan bool, runtime.NumCPU())

	jsonparser.ArrayEach(bodyText, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		wg.Add(1)
		sem <- true
		go func(value []byte) {
			defer wg.Done()
			defer func() { <-sem }()

			var rom Rom
			rom.Title, _ = jsonparser.GetString(value, "title")

			platformString, _ := jsonparser.GetString(value, "system")
			p, ok := VimmPlatforms[platformString]
			if ok {
				rom.Platform = p
			}

			rom.PageUrl, _ = jsonparser.GetString(value, "url")

			req, err := http.NewRequest("GET", rom.PageUrl, nil)
			req.Header.Set("accept", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}

			//Download Url
			doc.Find("#dl_form").EachWithBreak(func(i int, s *goquery.Selection) bool {
				if i > 0 {
					return false
				}
				mediaId, _ := s.Find("input").First().Attr("value")
				rom.DownloadUrl = "https://dl3.vimm.net/?mediaId=" + mediaId
				return false
			})

			doc.Find("canvas#canvas2").EachWithBreak(func(i int, s *goquery.Selection) bool {
				encodedRomName := s.AttrOr("data-v", "")

				if encodedRomName == "" {
					log.Println("No Rom Name found")
					return false
				}

				decodedRomName, err := base64.StdEncoding.DecodeString(encodedRomName)
				if err != nil {
					log.Println("Error decoding RomName:", err)
					return false
				}

				rom.RomName = string(decodedRomName)
				fmt.Println("RomTitle:", rom.RomName)

				return false
			})

			//Game Hashes
			doc.Find("#data-md5").EachWithBreak(func(i int, s *goquery.Selection) bool {
				if i > 0 {
					return false
				}

				rom.GameHash.MD5, err = s.Html()
				if err != nil {
					log.Println("Error getting HTML:", err)
				}

				return false
			})

			doc.Find("#data-sha1").EachWithBreak(func(i int, s *goquery.Selection) bool {
				if i > 0 {
					return false
				}

				rom.GameHash.SHA1, err = s.Html()
				if err != nil {
					log.Println("Error getting HTML:", err)
				}

				return false
			})

			doc.Find("#data-crc").EachWithBreak(func(i int, s *goquery.Selection) bool {
				if i > 0 {
					return false
				}

				rom.GameHash.CRC, err = s.Html()
				if err != nil {
					log.Println("Error getting HTML:", err)
				}

				return false
			})

			mu.Lock()
			roms = append(roms, rom)
			mu.Unlock()
		}(value)
	}, "games")

	wg.Wait()

	sort.Slice(roms, func(i, j int) bool {
		distI := levenshtein.ComputeDistance(roms[i].Title, gameName)
		distJ := levenshtein.ComputeDistance(roms[j].Title, gameName)
		return distI < distJ
	})

	return
}

func reqSetHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Referer", "https://vimm.net/vault/5172")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "counted=1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Priority", "u=0, i")
}

func DownloadGame(rom Rom, path string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rom.DownloadUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	reqSetHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(path, rom.Title+".7z"), bodyText, 0644)
	if err != nil {
		log.Fatal("Error writing file:", err)
		return
	}
}
