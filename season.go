package main

import "fmt"

func displaySeasonInfo(response Response, idSquad string) string {
	squad := response.Data[idSquad].Squads[0]
	seasonScoreInfo := squad.SeasonScoreInfo

	globalLeague := squad.GlobalLeagues[0]

	scoredPoints := formatNumberWithSpaces(seasonScoreInfo.Score)
	currentPlace := formatNumberWithSpaces(seasonScoreInfo.Place)
	totalPlaces := formatNumberWithSpaces(globalLeague.TotalPlaces)

	return fmt.Sprintf(
		"You scored %v points in the season and are in %vrd place out of %v \n",
		scoredPoints,
		currentPlace,
		totalPlaces,
	)
}
