package scraping

import (
	"bytes"
	"os"

	"strconv"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/buger/jsonparser"
	"github.com/joho/godotenv"

	"image"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Game struct {
	Title    string
	Id       int64
	Link     string
	CoverImg image.Image
	CoverUrl string
}

// TODO: Make it accept next page
// TODO: Add platform list fro Downloads\response_1755844081358.json
func TGDBGetGamesByName(searchName string) (games []Game) {
	client := &http.Client{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	apikey := os.Getenv("TGD_API_KEY")

	if apikey == "" {
		log.Println("no api found")
	}

	req, err := http.NewRequest("GET", "https://api.thegamesdb.net/v1/Games/ByGameName?apikey="+apikey+"&name="+url.QueryEscape(searchName), nil)
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

	urlCode, err := jsonparser.GetInt(bodyText, "code")
	if urlCode != 200 {
		log.Println("Fail to load games: ", urlCode)
	}

	if err != nil {
		log.Println(err)
	}

	jsonparser.ArrayEach(bodyText, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		gameName, err := jsonparser.GetString(value, "game_title")
		if err != nil {
			log.Println(err)
		}
		var newGame Game
		newGame.Title = gameName
		// gameNameList = append(gameNameList, gameName)
		gameId, err := jsonparser.GetInt(value, "id")
		if err != nil {
			log.Println(err)
		}
		// ids = append(ids, id)
		newGame.Id = gameId
		games = append(games, newGame)
	}, "data", "games")

	if err != nil {
		log.Printf("game_title not found: %v", err)
	}

	// id, err = jsonparser.GetInt(bodyText, "data", "games", "[0]", "id")
	if err != nil {
		log.Printf("id not found: %v", err)
	}

	return
}

func TGDBGetGameCover(url string) (cover image.Image) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:141.0) Gecko/20100101 Firefox/141.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "phpbb3_4vkdw_u=45897; phpbb3_4vkdw_k=j6b97uvgcjlt6lwc; phpbb3_4vkdw_sid=897a5c620aefb568fe7bb4a0ab8ad756")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	// req.Header.Set("If-Modified-Since", "Thu, 04 Aug 2022 12:58:15 GMT")
	// req.Header.Set("If-None-Match", `"62ebc267-d2aa"`)
	req.Header.Set("Priority", "u=0, i")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	// if resp.StatusCode != http.StatusOK {
	// 	log.Printlnf("Bad status: %d", resp.StatusCode)
	// }

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	cover, _, err = image.Decode(bytes.NewReader(bodyText))

	if err != nil {
		log.Println("Error Decoding image: ", err)
	}

	return
}

func TGDBGetGameCoverUrl(id int64) (url string) {

	client := &http.Client{}
	err := godotenv.Load()

	if err != nil {
		log.Printf("err loading: %v", err)
	}

	apikey := os.Getenv("TGD_API_KEY")

	url = "https://api.thegamesdb.net/v1/Games/Images?apikey=" + apikey + "&games_id=" + strconv.FormatInt(id, 10)
	req, err := http.NewRequest("GET", url, nil)
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

	filename, err := jsonparser.GetString(bodyText, "data", "images", strconv.FormatInt(id, 10), "[0]", "filename")
	if err != nil {
		log.Println("Error parsing Json:", err)
	}

	urlStart := "https://cdn.thegamesdb.net/images/original/"
	url = urlStart + filename
	return
}
