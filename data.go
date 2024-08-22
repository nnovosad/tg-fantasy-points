package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "golang.org/x/text/cases"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Data       map[string]DataInfo `json:"data"`
	Extensions struct {
		TransactionID string `json:"transactionID"`
	} `json:"extensions"`
}

type DataInfo struct {
	Squads []SquadInfo `json:"squads"`
}

type SquadInfo struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	CurrentTourInfo CurrentTourInfo `json:"currentTourInfo"`
	SeasonScoreInfo SeasonScoreInfo `json:"seasonScoreInfo"`
	GlobalLeagues   []GlobalLeagues `json:"globalLeagues"`
}

type CurrentTourInfo struct {
	Players   []PlayerInfo `json:"players"`
	ScoreInfo ScoreInfo    `json:"scoreInfo"`
	TourInfo  TourInfo     `json:"tour"`
}

type ScoreInfo struct {
	AverageScore float64 `json:"averageScore"`
	Score        int     `json:"score"`
}

type TourInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type PlayerInfo struct {
	IsCaptain     bool             `json:"isCaptain"`
	IsViceCaptain bool             `json:"isViceCaptain"`
	IsStarting    bool             `json:"isStarting"`
	SeasonPlayer  SeasonPlayerInfo `json:"seasonPlayer"`
	Score         int              `json:"score"`
}

type SeasonPlayerInfo struct {
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Role  string   `json:"role"`
	Team  TeamInfo `json:"team"`
}

type TeamInfo struct {
	Name string `json:"name"`
}

type SeasonScoreInfo struct {
	Place int `json:"place"`
	Score int `json:"score"`
}

type GlobalLeagues struct {
	Place       int `json:"place"`
	PlaceDiff   int `json:"placeDiff"`
	TotalPlaces int `json:"totalPlaces"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/points":
				handleCallback(update.Message.Chat.ID)
			default:
				fmt.Println("Unknown command")
			}
		}
	}
}

func handleCallback(chatID int64) {
	leagues := getLeagues()

	for country, data := range leagues {
		postBody, _ := json.Marshal(map[string]string{
			"query": data["query"],
		})

		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post("https://www.sports.ru/gql/graphql/", "application/json", responseBody)

		if err != nil {
			log.Fatalf("An Error Occurred %v", err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}

		idSquad := data["id"]

		country = uppercaseFirstCharacter(country)

		displayMatchesInfoMessage := scrapingTour(response, idSquad, data["tournament"])
		displayInfoTourMessage, displayResultsTourMessage := displayTourInfo(response, idSquad)
		displaySeasonInfoMessage := displaySeasonInfo(response, idSquad)

		output := "League: " + country + "\n" +
			"Tour information: " + displayInfoTourMessage + "\n" +
			"Matches information: \n" + displayMatchesInfoMessage + "\n" +
			"Tour results: " + displayResultsTourMessage + "\n" +
			"Season Information: " + displaySeasonInfoMessage

		telegramBotApiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")

		bot, err := tgbotapi.NewBotAPI(telegramBotApiToken)
		if err != nil {
			panic(err)
		}

		message := tgbotapi.NewMessage(chatID, output)
		_, err = bot.Send(message)
		if err != nil {
			log.Panic(err)
		}
	}
}

func getLeagues() map[string]map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	italyID := os.Getenv("ITALY_ID")
	italyTournament := os.Getenv("ITALY_TOURNAMENT")
	italyQuery := os.Getenv("ITALY_QUERY")

	russiaID := os.Getenv("RUSSIA_ID")
	russiaTournament := os.Getenv("RUSSIA_TOURNAMENT")
	russiaQuery := os.Getenv("RUSSIA_QUERY")

	englandID := os.Getenv("ENGLAND_ID")
	englandTournament := os.Getenv("ENGLAND_TOURNAMENT")
	englandQuery := os.Getenv("ENGLAND_QUERY")

	spainID := os.Getenv("SPAIN_ID")
	spainTournament := os.Getenv("SPAIN_TOURNAMENT")
	spainQuery := os.Getenv("SPAIN_QUERY")

	franceID := os.Getenv("FRANCE_ID")
	franceTournament := os.Getenv("FRANCE_TOURNAMENT")
	franceQuery := os.Getenv("FRANCE_QUERY")

	germanyID := os.Getenv("GERMANY_ID")
	germanyTournament := os.Getenv("GERMANY_TOURNAMENT")
	germanyQuery := os.Getenv("GERMANY_QUERY")

	portugalID := os.Getenv("PORTUGAL_ID")
	portugalTournament := os.Getenv("PORTUGAL_TOURNAMENT")
	portugalQuery := os.Getenv("PORTUGAL_QUERY")

	hollandID := os.Getenv("HOLLAND_ID")
	hollandTournament := os.Getenv("HOLLAND_TOURNAMENT")
	hollandQuery := os.Getenv("HOLLAND_QUERY")

	turkeyID := os.Getenv("TURKEY_ID")
	turkeyTournament := os.Getenv("TURKEY_TOURNAMENT")
	turkeyQuery := os.Getenv("TURKEY_QUERY")

	championshipID := os.Getenv("CHAMPIONSHIP_ID")
	championshipTournament := os.Getenv("CHAMPIONSHIP_TOURNAMENT")
	championshipQuery := os.Getenv("CHAMPIONSHIP_QUERY")

	leagues := map[string]map[string]string{
		"italy": {
			"id":         italyID,
			"tournament": italyTournament,
			"query":      italyQuery,
		},
		"russia": {
			"id":         russiaID,
			"tournament": russiaTournament,
			"query":      russiaQuery,
		},
		"england": {
			"id":         englandID,
			"tournament": englandTournament,
			"query":      englandQuery,
		},
		"spain": {
			"id":         spainID,
			"tournament": spainTournament,
			"query":      spainQuery,
		},
		"france": {
			"id":         franceID,
			"tournament": franceTournament,
			"query":      franceQuery,
		},
		"germany": {
			"id":         germanyID,
			"tournament": germanyTournament,
			"query":      germanyQuery,
		},
		"portugal": {
			"id":         portugalID,
			"tournament": portugalTournament,
			"query":      portugalQuery,
		},
		"holland": {
			"id":         hollandID,
			"tournament": hollandTournament,
			"query":      hollandQuery,
		},
		"turkey": {
			"id":         turkeyID,
			"tournament": turkeyTournament,
			"query":      turkeyQuery,
		},
		"championship": {
			"id":         championshipID,
			"tournament": championshipTournament,
			"query":      championshipQuery,
		},
	}

	return leagues
}
