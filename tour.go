package main

import "fmt"

func displayTourInfo(response Response) {
	squad := response.Data["id_105467854"].Squads[0]
	scoreInfo := squad.CurrentTourInfo.ScoreInfo

	fmt.Printf("You scored %v points. Average score %v points \n", scoreInfo.Score, scoreInfo.AverageScore)
}
