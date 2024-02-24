package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func scrapingTour(response Response) {
	c := colly.NewCollector()

	c.OnHTML("li.teaser-event", func(e *colly.HTMLElement) {
		match := e.ChildText(".teaser-event__board")
		status := e.ChildText(".teaser-event__status")

		squad := response.Data["id_105467854"].Squads[0]
		players := squad.CurrentTourInfo.Players

		preparedMatch := removeExtraSpaces(match)
		preparedStatus := getStatusMatch(status)

		serialNumber := e.Index
		serialNumber++

		fmt.Printf("%v. %v %v \n", serialNumber, preparedStatus, preparedMatch)

		if preparedStatus == "завершен" {
			for _, v := range players {
				teamName := v.SeasonPlayer.Team.Name
				if strings.Contains(match, teamName) {
					fmt.Printf("--- %v scored %v points \n", v.SeasonPlayer.Name, v.Score)
				}
			}
		}
	})

	err := c.Visit("https://www.sports.ru/seria-a/")
	if err != nil {
		return
	}
}

func removeExtraSpaces(input string) string {
	re := regexp.MustCompile(`\s+`)
	stripped := re.ReplaceAllString(input, " ")

	stripped = strings.TrimSpace(stripped)

	return stripped
}

func getStatusMatch(statusMatch string) string {
	regexpPattern := regexp.MustCompile(`(\d{2}\.\d{2} \d{2}:\d{2})|(первый тайм|второй тайм|завершен)`)

	match := regexpPattern.FindStringIndex(statusMatch)
	if match != nil {
		statusMatch = statusMatch[:match[1]]
	}

	return statusMatch
}
