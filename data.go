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

	commonID := os.Getenv("QUERY_COMMON_ID")

	italyID := os.Getenv("ITALY_ID")
	italyTournament := os.Getenv("ITALY_TOURNAMENT")
	italyFantasyQueriesID := os.Getenv("ITALY_FANTASY_QUERIES_ID")
	italyOddsQueriesID := os.Getenv("ITALY_ODDS_QUERIES_ID")
	italyName := os.Getenv("ITALY_NAME")
	italySquadID := os.Getenv("ITALY_SQUAD_ID")
	italyQuery := prepareBodyRequest(commonID, italyID, italyFantasyQueriesID, italyOddsQueriesID, italyName, italySquadID)

	russiaID := os.Getenv("RUSSIA_ID")
	russiaTournament := os.Getenv("RUSSIA_TOURNAMENT")
	russiaFantasyQueriesID := os.Getenv("RUSSIA_FANTASY_QUERIES_ID")
	russiaOddsQueriesID := os.Getenv("RUSSIA_ODDS_QUERIES_ID")
	russiaName := os.Getenv("RUSSIA_NAME")
	russiaSquadID := os.Getenv("RUSSIA_SQUAD_ID")
	russiaQuery := prepareBodyRequest(commonID, russiaID, russiaFantasyQueriesID, russiaOddsQueriesID, russiaName, russiaSquadID)

	englandID := os.Getenv("ENGLAND_ID")
	englandTournament := os.Getenv("ENGLAND_TOURNAMENT")
	englandFantasyQueriesID := os.Getenv("ENGLAND_FANTASY_QUERIES_ID")
	englandOddsQueriesID := os.Getenv("ENGLAND_ODDS_QUERIES_ID")
	englandName := os.Getenv("ENGLAND_NAME")
	englandSquadID := os.Getenv("ENGLAND_SQUAD_ID")
	englandQuery := prepareBodyRequest(commonID, englandID, englandFantasyQueriesID, englandOddsQueriesID, englandName, englandSquadID)

	spainID := os.Getenv("SPAIN_ID")
	spainTournament := os.Getenv("SPAIN_TOURNAMENT")
	spainFantasyQueriesID := os.Getenv("SPAIN_FANTASY_QUERIES_ID")
	spainOddsQueriesID := os.Getenv("SPAIN_ODDS_QUERIES_ID")
	spainName := os.Getenv("SPAIN_NAME")
	spainSquadID := os.Getenv("SPAIN_SQUAD_ID")
	spainQuery := prepareBodyRequest(commonID, spainID, spainFantasyQueriesID, spainOddsQueriesID, spainName, spainSquadID)

	franceID := os.Getenv("FRANCE_ID")
	franceTournament := os.Getenv("FRANCE_TOURNAMENT")
	franceFantasyQueriesID := os.Getenv("FRANCE_FANTASY_QUERIES_ID")
	franceOddsQueriesID := os.Getenv("FRANCE_ODDS_QUERIES_ID")
	franceName := os.Getenv("FRANCE_NAME")
	franceSquadID := os.Getenv("FRANCE_SQUAD_ID")
	franceQuery := prepareBodyRequest(commonID, franceID, franceFantasyQueriesID, franceOddsQueriesID, franceName, franceSquadID)

	germanyID := os.Getenv("GERMANY_ID")
	germanyTournament := os.Getenv("GERMANY_TOURNAMENT")
	germanyFantasyQueriesID := os.Getenv("GERMANY_FANTASY_QUERIES_ID")
	germanyOddsQueriesID := os.Getenv("GERMANY_ODDS_QUERIES_ID")
	germanyName := os.Getenv("GERMANY_NAME")
	germanySquadID := os.Getenv("GERMANY_SQUAD_ID")
	germanyQuery := prepareBodyRequest(commonID, germanyID, germanyFantasyQueriesID, germanyOddsQueriesID, germanyName, germanySquadID)

	portugalID := os.Getenv("PORTUGAL_ID")
	portugalTournament := os.Getenv("PORTUGAL_TOURNAMENT")
	portugalFantasyQueriesID := os.Getenv("PORTUGAL_FANTASY_QUERIES_ID")
	portugalOddsQueriesID := os.Getenv("PORTUGAL_ODDS_QUERIES_ID")
	portugalName := os.Getenv("PORTUGAL_NAME")
	portugalSquadID := os.Getenv("PORTUGAL_SQUAD_ID")
	portugalQuery := prepareBodyRequest(commonID, portugalID, portugalFantasyQueriesID, portugalOddsQueriesID, portugalName, portugalSquadID)

	hollandID := os.Getenv("HOLLAND_ID")
	hollandTournament := os.Getenv("HOLLAND_TOURNAMENT")
	hollandFantasyQueriesID := os.Getenv("HOLLAND_FANTASY_QUERIES_ID")
	hollandOddsQueriesID := os.Getenv("HOLLAND_ODDS_QUERIES_ID")
	hollandName := os.Getenv("HOLLAND_NAME")
	hollandSquadID := os.Getenv("HOLLAND_SQUAD_ID")
	hollandQuery := prepareBodyRequest(commonID, hollandID, hollandFantasyQueriesID, hollandOddsQueriesID, hollandName, hollandSquadID)

	turkeyID := os.Getenv("TURKEY_ID")
	turkeyTournament := os.Getenv("TURKEY_TOURNAMENT")
	turkeyFantasyQueriesID := os.Getenv("TURKEY_FANTASY_QUERIES_ID")
	turkeyOddsQueriesID := os.Getenv("TURKEY_ODDS_QUERIES_ID")
	turkeyName := os.Getenv("TURKEY_NAME")
	turkeySquadID := os.Getenv("TURKEY_SQUAD_ID")
	turkeyQuery := prepareBodyRequest(commonID, turkeyID, turkeyFantasyQueriesID, turkeyOddsQueriesID, turkeyName, turkeySquadID)

	championshipID := os.Getenv("CHAMPIONSHIP_ID")
	championshipTournament := os.Getenv("CHAMPIONSHIP_TOURNAMENT")
	championshipFantasyQueriesID := os.Getenv("CHAMPIONSHIP_FANTASY_QUERIES_ID")
	championshipOddsQueriesID := os.Getenv("CHAMPIONSHIP_ODDS_QUERIES_ID")
	championshipName := os.Getenv("CHAMPIONSHIP_NAME")
	championshipSquadID := os.Getenv("CHAMPIONSHIP_SQUAD_ID")
	championshipQuery := prepareBodyRequest(commonID, championshipID, championshipFantasyQueriesID, championshipOddsQueriesID, championshipName, championshipSquadID)

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

func prepareBodyRequest(commonId string, tournamentId string, queryId string, oddsId string, tournamentName string, squadId string) string {
	return fmt.Sprintf(
		"{ %s: userQueries { current { user { id nick url picture(input: { ext: PNG, resize: SIZE_64_64 }) { url } } commentsLinksEnabled } } %s: fantasyQueries { tournament(id: \"%s\", source: HRU) { currentSeason { id currentSquad { id } } } } %s: oddsQueries { bookerByPlacement(input: { placementName: \"FANTASY_MATCH_ITALY\" iso2Country: \"BY\" }) { primaryColor secondaryColor CTA CTALink popupBanner popupCTALink } } %s: fantasyQueries { squads(input: {squadID: \"%s\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds( placementName: \"FANTASY_MATCH_ITALY\" iso2Country: \"BY\" ) { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description form } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_223213878: oddsQueries { bookerByPlacement(input: { placementName: \"MENU_WEB_LEFT\", iso2Country: \"BY\" }) { logoUrl primaryColor secondaryColor name title lead url CTA popupText } } id_256551909: oddsQueries { bookerByPlacement(input: { placementName: \"MENU_WEB_RIGHT\", iso2Country: \"BY\" }) { logoUrl primaryColor secondaryColor name title lead url CTA popupText } } id_223213878: oddsQueries { bookerByPlacement(input: { placementName: \"MENU_WEB_LEFT\", iso2Country: \"BY\" }) { logoUrl primaryColor secondaryColor name title lead url CTA popupText } } id_123679826: commentQueries { list( objectId: \"ru_fantasy_italy\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
		commonId, queryId, tournamentName, oddsId, tournamentId, squadId,
	)
}
