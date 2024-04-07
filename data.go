package main

import (
	"bytes"
	"encoding/json"
	_ "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
}

type ScoreInfo struct {
	AverageScore float64 `json:"averageScore"`
	Score        int     `json:"score"`
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
	leagues := map[string]map[string]string{
		"italy": {
			"id":         "id_105467854",
			"tournament": "seria-a",
			"query":      "{ id_266299461: fantasyQueries { tournament(id: \"italy\", source: HRU) { currentSeason { id currentSquad { id } } } } id_105467854: fantasyQueries { squads(input: {squadID: \"197100\"}) { id name user { id nick url } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_ITALY\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_123679826: commentQueries { list( objectId: \"ru_fantasy_italy\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url } parentComment { id text published { epoch } author { id nick url } } } pagination { hasNextPage } } } }",
		},
		"russia": {
			"id":         "id_219482870",
			"tournament": "rfpl",
			"query":      "{ id_14272865: fantasyQueries { tournament(id: \"russia\", source: HRU) { currentSeason { id currentSquad { id } } } } id_219482870: fantasyQueries { squads(input: {squadID: \"156736\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_RUSSIA\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description form } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_31268918: commentQueries { list( objectId: \"ru_fantasy_russia\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
		},
		"england": {
			"id":         "id_8550644",
			"tournament": "premier-league",
			"query":      "{ id_5639147: fantasyQueries { tournament(id: \"england\", source: HRU) { currentSeason { id currentSquad { id } } } } id_8550644: fantasyQueries { squads(input: {squadID: \"178794\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_ENGLAND\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description form } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_43471852: commentQueries { list( objectId: \"ru_fantasy_england\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
		},
		"spain": {
			"id":         "id_4250943",
			"tournament": "la-liga",
			"query":      "{ id_93213035: fantasyQueries { tournament(id: \"spain\", source: HRU) { currentSeason { id currentSquad { id } } } } id_4250943: fantasyQueries { squads(input: {squadID: \"178816\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_SPAIN\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description form } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_180637292: commentQueries { list( objectId: \"ru_fantasy_spain\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
		},
		"france": {
			"id":         "id_263542169",
			"tournament": "ligue-1",
			"query":      "{ id_29235521: fantasyQueries { tournament(id: \"france\", source: HRU) { currentSeason { id currentSquad { id } } } } id_263542169: fantasyQueries { squads(input: {squadID: \"178821\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_FRANCE\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description form } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_242963542: commentQueries { list( objectId: \"ru_fantasy_france\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
		},
		"germany": {
			"id":         "id_242532932",
			"tournament": "bundesliga",
			"query":      "{ id_40766421: fantasyQueries { tournament(id: \"germany\", source: HRU) { currentSeason { id currentSquad { id } } } } id_242532932: fantasyQueries { squads(input: {squadID: \"195113\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_GERMANY\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description form } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_82046658: commentQueries { list( objectId: \"ru_fantasy_germany\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
		},
	}

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

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}

		idSquad := data["id"]

		preparedCountry := strings.Title(country)

		displayMatchesInfoMessage := scrapingTour(response, idSquad, data["tournament"])
		displayTourInfoMessage := displayTourInfo(response, idSquad)
		displaySeasonInfoMessage := displaySeasonInfo(response, idSquad)

		output := "League: " + preparedCountry + "\n" +
			"Matches information: \n" + displayMatchesInfoMessage + "\n" +
			"Tour information: " + displayTourInfoMessage + "\n" +
			"Season Information: " + displaySeasonInfoMessage

		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		telegramBotApiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")

		bot, err := tgbotapi.NewBotAPI(telegramBotApiToken)
		if err != nil {
			panic(err)
		}

		message := tgbotapi.NewMessage(514411911, output)
		_, err = bot.Send(message)
		if err != nil {
			log.Panic(err)
		}

		//fmt.Print("Start display team \n")
		//displayTeam(response, idSquad)
		//fmt.Print("Finish display team \n")
	}
}
