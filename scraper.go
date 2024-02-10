package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func scrapingTour() {
	c := colly.NewCollector()

	c.OnHTML("li.teaser-event", func(e *colly.HTMLElement) {
		match := e.ChildText(".teaser-event__board")
		status := e.ChildText(".teaser-event__status")

		preparedMatch := removeExtraSpaces(match)
		preparedStatus := getStatusMatch(status)

		serialNumber := e.Index
		serialNumber++

		fmt.Printf("%v. %v %v \n", serialNumber, preparedStatus, preparedMatch)
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
