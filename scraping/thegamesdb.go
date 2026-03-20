package scraping

import (
	"os"

	"strconv"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/buger/jsonparser"
	"github.com/joho/godotenv"

	"io"
	"log"
	"net/http"
	"net/url"
)

type Game struct {
	Title    string
	Id       int64
	Link     string
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
