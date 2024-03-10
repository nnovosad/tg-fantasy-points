package main

import "fmt"

func displaySeasonInfo(response Response) {
	squad := response.Data["id_105467854"].Squads[0]
	seasonScoreInfo := squad.SeasonScoreInfo

	globalLeague := squad.GlobalLeagues[0]

	fmt.Printf(
		"You scored %v points in the season and are in %vrd place out of %v \n",
		seasonScoreInfo.Score,
		seasonScoreInfo.Place,
		globalLeague.TotalPlaces,
	)
}
