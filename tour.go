package main

import "fmt"

func displayTourInfo(response Response, idSquad string) {
	squad := response.Data[idSquad].Squads[0]
	scoreInfo := squad.CurrentTourInfo.ScoreInfo

	fmt.Printf("You scored %v points. Average score %v points \n", scoreInfo.Score, scoreInfo.AverageScore)
}
