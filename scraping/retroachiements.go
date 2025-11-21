package scraping

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/joho/godotenv"
)

// Retroachievements Supported Systems
//
//
// | System / Platform                       | Hash Type          |
// | --------------------------------------- | ------------------ |
// | Atari 2600                              | MD5                |
// | Atari 7800                              | MD5 (header-aware) |
// | Atari Jaguar                            | MD5                |
// | Atari Jaguar CD                         | Custom             |
// | Atari Lynx                              | MD5 (header-aware) |
// | Bandai WonderSwan / Color               | MD5                |
// | Channel F                               | MD5                |
// | ColecoVision                            | MD5                |
// | Dreamcast                               | Custom             |
// | Famicom Disk System                     | MD5 (header-aware) |
// | Game Boy / Color / Advance              | MD5                |
// | Game Gear                               | MD5                |
// | GameCube                                | Custom             |
// | Intellivision                           | MD5                |
// | MSX / MSX2                              | MD5                |
// | Master System                           | MD5                |
// | Mega Drive / Genesis                    | MD5                |
// | NEC PC-8001 / 8801                      | MD5                |
// | NES / Famicom                           | Custom             |
// | Neo Geo CD                              | Custom             |
// | Neo Geo Pocket / Color                  | MD5                |
// | Nintendo 64                             | MD5                |
// | Nintendo DS                             | Custom             |
// | Odyssey²                                | MD5                |
// | PC Engine / TurboGrafx / SuperGrafx     | MD5 (header-aware) |
// | PC Engine CD / TurboGrafx-CD            | Custom             |
// | PC-FX                                   | Custom             |
// | PSP                                     | Custom             |
// | PlayStation                             | Custom             |
// | PlayStation 2                           | Custom             |
// | Pokémon Mini                            | MD5                |
// | SG-1000                                 | MD5                |
// | SNES / SFC / Satellaview / Sufami Turbo | MD5 (header-aware) |
// | Saturn                                  | Custom             |
// | Sega 32X                                | MD5                |
// | Sega CD                                 | Custom             |
// | Vectrex                                 | MD5                |
// | Virtual Boy                             | MD5                |
// | WASM-4                                  | MD5                |
// | Watara Supervision                      | MD5                |
// | Wellback Mega Duck                      | MD5                |

var RAconsoles = map[string]int{
	"Atari 2600":     25, //MD5
	"Atari 5200":     50, //?
	"Atari 7800":     51, //MD5
	"CD-i":           42, //No achiements
	"Dreamcast":      40, //Custom
	"GBA":            5,  //MD5
	"Game Boy Color": 6,  //MD5
	"Game Boy":       4,  //MD5
	"Game Gear":      15, //MD5
	"GameCube":       16, //Custom
	"Genesis":        1,  //MD5
	"Jaguar CD":      77, //Custom
	"Jaguar":         17, //MD5
	"Lynx":           13, //MD5
	"Master System":  11, //MD5
	"Nintendo 3DS":   62, //No achiements
	"Nintendo 64":    2,  //MD5
	"Nintendo DS":    18, //Custom
	"Nintendo":       62, //Custom
	"PS Portable":    41, //Custom
	"PlayStation 2":  21, //Custom
	"PlayStation":    41, //Custom
	"Saturn":         39, //Custom
	"Sega CD":        9,  //Custom
	"Super Nintendo": 3,  //MD5
	"TurboGrafx-16":  8,  //MD5
	"TurboGrafx-CD":  76, //Custom
	"Virtual Boy":    28, //MD5j
	"Wii":            19, //No achiements
	"Xbox":           22, //No achiements
	// "PlayStation 3" :
	// "Sega 32X" :
	// "WiiWare" :
	// "Xbox 360" :
	// Xbox 360 (Digital)
}

func normalize(s string) string {
	re := regexp.MustCompile(`[^a-z]+`)
	s = strings.ToLower(s)
	s = re.ReplaceAllString(s, "")
	return s
}

func structHasValue(v any, target any) bool {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if reflect.DeepEqual(field.Interface(), target) {
			return true
		}
	}
	return false
}

func RAvalidateHash(rom Rom) (isValid bool) {

	client := &http.Client{}

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading .env: %v", err)
	}

	apikey := os.Getenv("ACHIEVEMENTS_API_KEY")
	if apikey == "" {
		log.Println("no API key found")
		return
	}

	baseURL := "https://retroachievements.org/API/API_GetGameList.php"

	v := url.Values{}
	v.Add("y", apikey)
	log.Println(rom.Platform)
	v.Add("i", strconv.Itoa(RAconsoles[rom.Platform]))
	v.Add("h", "1")
	v.Add("f", "1")

	reqURL := fmt.Sprintf("%s?%s", baseURL, v.Encode())
	log.Println(reqURL)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Printf("error creating request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
		return
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// bodyText, err := os.ReadFile("./test2.json")
	if err != nil {
		panic(err)
	}

	isValid = false

	jsonparser.ArrayEach(bodyText, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		raTitle, err := jsonparser.GetString(value, "Title")
		if err != nil {
			return
		}

		if normalize(raTitle) != normalize(rom.Title) {
			return
		}

		jsonparser.ArrayEach(value, func(hashValue []byte, dataType jsonparser.ValueType, offset int, err error) {
			if structHasValue(rom.GameHash, string(hashValue)) {
				isValid = true
			}
		}, "Hashes")
	})
	return
}
