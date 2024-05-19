package main

import (
	"fmt"
	"strings"
)

func displayTourInfo(response Response, idSquad string) (string, string) {
	squad := response.Data[idSquad].Squads[0]
	scoreInfo := squad.CurrentTourInfo.ScoreInfo
	tourInfo := squad.CurrentTourInfo.TourInfo

	tourName := tourInfo.Name
	tourStatus := tourInfo.Status

	tourStatus = uppercaseFirstCharacter(tourStatus)

	tourStatus = strings.Replace(tourStatus, "_", " ", -1)

	return fmt.Sprintf("%v. %v", tourName, tourStatus),
		fmt.Sprintf("You scored %v points. Average score %v points \n", scoreInfo.Score, scoreInfo.AverageScore)
}
