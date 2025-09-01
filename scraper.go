package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func scrapingTour(response Response, idSquad string, tournament string) string {
	c := colly.NewCollector()

	var output = ""

	c.OnHTML("li.teaser-event", func(e *colly.HTMLElement) {
		match := e.ChildText(".teaser-event__board")
		status := e.ChildText(".teaser-event__status")

		squad := response.Data[idSquad].Squads[0]
		players := squad.CurrentTourInfo.Players

		preparedMatch := removeExtraSpaces(match)
		preparedStatus := getStatusMatch(status)

		serialNumber := e.Index
		serialNumber++

		output += fmt.Sprintf("%v. %v %v \n", serialNumber, preparedStatus, preparedMatch)

		output += printPlayerInfo(players, match, preparedStatus)
	})

	err := c.Visit("https://www.sports.ru/football/tournament/" + tournament)
	if err != nil {
		return err.Error()
	}

	if output == "" {
		output = "No matches found" + "\n"
	}

	return output
}

func removeExtraSpaces(input string) string {
	re := regexp.MustCompile(`\s+`)
	stripped := re.ReplaceAllString(input, " ")

	stripped = strings.TrimSpace(stripped)

	return stripped
}

func getStatusMatch(statusMatch string) string {
	regexpPattern := regexp.MustCompile(`(\d{2}\.\d{2} \d{2}:\d{2})|(первый тайм|второй тайм|завершен|перерыв)`)

	match := regexpPattern.FindStringIndex(statusMatch)
	if match != nil {
		statusMatch = statusMatch[:match[1]]
	}

	return statusMatch
}

func printPlayerInfo(players PlayersSlice, match string, statusMatch string) string {
	var output = ""

	for _, v := range players {
		playerTeamName := v.SeasonPlayer.Team.Name

		if isPlayerTeam(match, playerTeamName) {
			playerStatus := "Main cast"

			role := uppercaseFirstCharacter(v.SeasonPlayer.Role)

			if !v.IsStarting {
				playerStatus = "On the bench"
			}

			if statusMatch == "завершен" {
				output += fmt.Sprintf(
					"--- %v(%v) scored %v points. %v. %v \n",
					v.SeasonPlayer.Name, playerTeamName, v.Score, playerStatus, role)
			} else {
				output += fmt.Sprintf(
					"--- Can play %v(%v). %v. %v \n",
					v.SeasonPlayer.Name, playerTeamName, playerStatus, role)
			}
		}
	}

	return output
}
