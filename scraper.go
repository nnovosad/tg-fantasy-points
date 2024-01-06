package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func scraping() {
	c := colly.NewCollector()

	c.OnHTML("li.teaser-event", func(e *colly.HTMLElement) {
		matches := e.ChildText(".teaser-event__board")
		status := e.ChildText(".teaser-event__status")
		href := e.ChildAttr(".teaser-event__board-score", "href")
		fullUrl := "https://www.sports.ru" + href

		fmt.Printf("Match %v %v: %v \n", matches, status, fullUrl)
	})

	err := c.Visit("https://www.sports.ru/seria-a/")
	if err != nil {
		return
	}
}
