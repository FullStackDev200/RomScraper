package scraping

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/PuerkitoBio/goquery"
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
	Platform    string
	GameHash    GameHash
	CoverUrl    string
	DownloadUrl string
	PageUrl     string
}

func VimmSearchRoms(gameName string, filter string) (roms []Rom) {
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
			rom.Platform, _ = jsonparser.GetString(value, "system")
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

			doc.Find("#dl_form").Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					return
				}
				mediaId, _ := s.Find("input").First().Attr("value")
				rom.DownloadUrl = "https://dl3.vimm.net/?mediaId=" + mediaId
			})

			doc.Find("img[title=\"Click to enlarge\"]").Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					return
				}
				rom.CoverUrl, _ = s.First().Attr("src")
			})

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
	return
}

func DownloadGame(rom Rom, path string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rom.DownloadUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

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
