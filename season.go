package main

import "fmt"

func displaySeasonInfo(response Response, idSquad string) string {
	squad := response.Data[idSquad].Squads[0]
	seasonScoreInfo := squad.SeasonScoreInfo

	globalLeague := squad.GlobalLeagues[0]

	seasonScoreInfoPlace := seasonScoreInfo.Place
	globalLeagueTotalPlaces := globalLeague.TotalPlaces

	scoredPoints := formatNumberWithSpaces(seasonScoreInfo.Score)
	currentPlace := formatNumberWithSpaces(seasonScoreInfoPlace)
	totalPlaces := formatNumberWithSpaces(globalLeagueTotalPlaces)

	percentileRank := prepareRank(seasonScoreInfoPlace, globalLeagueTotalPlaces)

	return fmt.Sprintf(
		"You scored %v points in the season and are in %vrd place out of %v. Rank: %v \n",
		scoredPoints,
		currentPlace,
		totalPlaces,
		percentileRank,
	)
}
