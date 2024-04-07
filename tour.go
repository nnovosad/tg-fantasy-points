package main

import "fmt"

func displayTourInfo(response Response, idSquad string) string {
	squad := response.Data[idSquad].Squads[0]
	scoreInfo := squad.CurrentTourInfo.ScoreInfo

	return fmt.Sprintf("You scored %v points. Average score %v points \n", scoreInfo.Score, scoreInfo.AverageScore)
}
