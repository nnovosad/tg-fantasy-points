package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type FantasyData struct {
	Data string
}

func main() {
	postBody, _ := json.Marshal(map[string]string{
		"query": "{ id_266299461: fantasyQueries { tournament(id: \"italy\", source: HRU) { currentSeason { id currentSquad { id } } } } id_105467854: fantasyQueries { squads(input: {squadID: \"197100\"}) { id name user { id nick url } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_ITALY\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_123679826: commentQueries { list( objectId: \"ru_fantasy_italy\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url } parentComment { id text published { epoch } author { id nick url } } } pagination { hasNextPage } } } }",
		//"query": "{ id_266299461: fantasyQueries { tournament(id: \"italy\", source: HRU) { currentSeason { id currentSquad { id } } } } id_105467854: fantasyQueries { squads(input: {squadID: \"197100\"}) { id name user { id nick url picture(input: { resize: SIZE_64_64 } ) { url } retina: picture(input: { resize: SIZE_128_128 } ) { url } } currentTourInfo { isNotLimit tour { name status id transfersFinishedAt constraints { totalTransfers maxSameTeamPlayers } season { info { playerPrices teams { id name } constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } matches { id matchStatus scheduledAtStamp dateOnly links { sportsRu } home { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } away { team { name logo { desktop: resize(width: \"60\", height: \"60\") desktop__2x: resize(width: \"120\", height: \"120\") mobile: resize(width: \"60\", height: \"60\") mobile__2x: resize(width: \"120\", height: \"120\") original: main } lastFive { result pointsDiff match { id links { link } scheduledAt home { team { id name } score } away { team { id name } score } } } } score } bettingOdds(placementName: \"FANTASY_MATCH_ITALY\") { outcomes: line1x2 { home: h draw: x away: a } bookmaker { id name primaryColor secondaryColor url } } } } scoreInfo { place score totalPlaces averageScore } totalPrice currentBalance transfersLeft topPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } topTransferPlayers{ id name price statObject { name firstName lastName } seasonScoreInfo { score } team { name svgKit { url } } } players { isCaptain isViceCaptain isStarting isPointsCount substitutePriority statDetails { score reason } seasonPlayer { id name price role statObject { name firstName lastName links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } team { id name svgKit { url } statObject { name links { sportsRu } desktop: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } desktop__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } mobile: logotype(input: {resize: SIZE_128_128, multi: X1}) { url } mobile__2x: logotype(input: {resize: SIZE_128_128, multi: X2}) { url } original: logotype(input: {resize: ORIGINAL, multi: X1}) { url } } } seasonScoreInfo { place score totalPlaces averageScore } gameStat { goals assists goalsConceded yellowCards redCards fieldMinutes saves } status { status description } } statDetails{ score reason } statPlayer { goals assists yellowCards redCards goalsConceded } score playedMatches playedMatchesTour } } seasonScoreInfo { score place } globalLeagues: leagues(input:{type: GENERAL}){ league { id name type } place totalPlaces placeDiff } regionalLeagues: leagues(input:{type: REGIONAL}){ league{ id name type } place totalPlaces placeDiff } season { id isActive tournament { id name } tours { id status name startedAt } info { teams { id name statObject { logo { desktop: resize(width: \"100\", height: \"100\") desktop__2x: resize(width: \"200\", height: \"200\") mobile: resize(width: \"100\", height: \"100\") mobile__2x: resize(width: \"200\", height: \"200\") original: main } } } playerPrices constraints { fullRoster { role minCount maxCount row } startingRoster { role minCount maxCount row } } } } } } id_123679826: commentQueries { list( objectId: \"ru_fantasy_italy\" objectClass: CHAT order: NEW first: 20 after: \"undefined\" ) { comments { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } parentComment { id text published { epoch } author { id nick url picture(input: {resize: SIZE_32_32}) { url } } } } pagination { hasNextPage } } } }",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://www.sports.ru/gql/graphql/", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	log.Printf(sb)

	//data := FantasyData{}
	//
	//jsonErr := json.Unmarshal([]byte(sb), &data)
	//if jsonErr != nil {
	//	log.Fatal(jsonErr)
	//}

	//log.Println(data)
}
