package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func HandleLevelsScrape() {
	// Set up a new colly collector
	c := colly.NewCollector()

	// Set up a callback function to be called for each HTML response
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		// Find the cells in the row
		cells := e.ChildTexts("td")

		// Skip rows that don't have enough cells
		if len(cells) < 3 {
			return
		}
		// Extract the company name and pay range from the cells
		companyName := cells[0]
		payRange := cells[1]

		// Print the company name and pay range
		fmt.Println(companyName + ": " + payRange)
	})

	// Set up a callback function to be called for each request
	c.OnRequest(func(r *colly.Request) {
		// Print the URL of the request
		fmt.Println("Visiting", r.URL)
	})

	// Set up a callback function to be called for each response error
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	// Set up a callback function to be called after all requests are completed
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Send a request to the first page of results
	c.Visit("https://www.levels.fyi/t/software-engineer?countryId=254&country=254&gender=male&yoeChoice=mid&sinceDate=year&limit=50&sortBy=total_compensation&sortOrder=DESC")

	// Wait until all requests are completed
	c.Wait()
}
