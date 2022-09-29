package main

import (
	"fmt"
	"mcampbellr/webscrapper/notificator"

	"github.com/gocolly/colly"
	"github.com/robfig/cron"
)

func main() {
	crn := cron.New()
	crn.AddFunc("@every 5s", func() { scrapper(crn) })
	crn.Start()
	fmt.Scanln()
}

func scrapper(c *cron.Cron) {
	fmt.Println("Running scrapper")
	cl := colly.NewCollector()

	cl.OnHTML(".fg-inner-page-nav__list", func(e *colly.HTMLElement) {
		isSold := false

		e.ForEach("a", func(i int, h *colly.HTMLElement) {
			linkText := h.Text

			if linkText == "Sold out t" {
				isSold = true
			}
		})

		if !isSold {
			notificator.NotifyIfNotSoldOut()
			c.Stop()
		}
	})

	cl.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	cl.Visit("https://kinesis-ergo.com/keyboards/advantage360/")
}
