package scraping

import (
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{}

func GetPossibleGames(gameName string) (possibleGames []Game) {
	searchURL := "https://www.emulatorgames.net/search/?kw=" + url.QueryEscape(gameName)

	resp := ReqEmulatorgames(searchURL)
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".site-list > li").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())

		romLink, _ := s.Find("[href]").Attr("href")
		link := regexp.MustCompile(`/roms/[^/]+/`).ReplaceAllString(romLink, `/download/?rom=`)
		link = strings.TrimSuffix(link, "/")

		imgLink, exists := s.Find("[src]").Attr("src")
		if !exists || imgLink == "" {
			imgLink = ""
		}

		var coverImg image.Image

		if imgLink != "" {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", imgLink, nil)
			if err != nil {
				log.Println("Image request error:", err)
			} else {

				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:140.0) Gecko/20100101 Firefox/140.0")
				req.Header.Set("Accept", "image/png,image/svg+xml,image/*;q=0.8,*/*;q=0.5")
				req.Header.Set("Accept-Language", "en-US,en;q=0.5")
				req.Header.Set("Referer", "https://www.emulatorgames.net/")

				coverResp, err := client.Do(req)
				if err != nil {
					log.Println("Image download failed:", err)
				} else {
					defer coverResp.Body.Close()

					log.Printf("Fetching image: %s", imgLink)

					img, format, err := image.Decode(coverResp.Body)
					if err != nil {
						log.Println("Image decode failed:", err, "Format given:", format)
					} else {
						coverImg = img
					}
				}
			}
		}

		possibleGames = append(possibleGames, Game{
			Title:    title,
			Link:     link,
			CoverImg: coverImg,
		})
	})

	return possibleGames
}

func ReqEmulatorgames(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request creation error:", err)
		return nil
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:140.0) Gecko/20100101 Firefox/140.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.emulatorgames.net/")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download %s. Status: %s\n", url, resp.Status)
		resp.Body.Close()
		return nil
	}

	return resp
}

func ExtractDataID(resp *http.Response) (int, error) {

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Select the download button that contains the data-id
	idStr, exists := doc.Find("[data-id]").Attr("data-id")
	if !exists {
		return 0, fmt.Errorf("data-id not found")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid data-id: %w", err)
	}

	return id, err
}

func ExtractImgUrl(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Find <img> tag with both classes: site-post-img and shadow-sm
	imgSrc, exists := doc.Find("img.site-post-img.shadow-sm").Attr("src")
	if !exists {
		return "", fmt.Errorf("image src not found")
	}

	return imgSrc, nil
}

func DownloadCover(url, gameName string) {
	// client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:140.0) Gecko/20100101 Firefox/140.0")
	req.Header.Set("Accept", "image/png,image/svg+xml,image/*;q=0.8,*/*;q=0.5")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://www.emulatorgames.net/")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Bad status code: %d", resp.StatusCode)
	}

	out, err := os.Create(gameName + ".jpg")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal("Error saving image:", err)
	}
}

func GetGameFromemulatorgames(url string, dataID int) {
	client := &http.Client{}

	var data = strings.NewReader(`get_type=post^&get_id=` + strconv.Itoa(dataID))
	req, err := http.NewRequest("POST", "https://www.emulatorgames.net/prompt/", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:140.0) Gecko/20100101 Firefox/140.0")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "https://www.emulatorgames.net")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", url)
	req.Header.Set("Cookie", "cf_clearance=6Hmnir8MafoOJcZXaEY7t2qT4U9OCgu8LQKjQWcy_bQ-1752159470-1.2.1.1-qwK.6kIrpXE1MziXCg3Nz3VhlzxLI0Qd_IB0RuV9Bbm_XTkwU6uqEO3BNbcm1_NlNgGzBVHmWhebBZWNWu0EeaGAx8ea9cQa5vkc3Ih3btq.QDpn9TAEzZU2VBmGNnQawss8cWveXlm332rJTLzy.U2VI1f9hq3SPE1Vrw01DcPRqAfev2tRuCvlT1hMv6pd9SBZXaMAx9Z6Bz6LBz_BcEllumDsVE.h4bzD2YB4O_Q")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

type Game struct {
	Title    string
	Id       int64
	Link     string
	CoverImg image.Image
}
