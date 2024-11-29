package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func amazonScrape() {
	// Initialize the Colly collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(), // Allow revisiting pages
		colly.Async(true),       // Make it asynchronous
	)

	// Handle request logging
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Handle product data extraction
	c.OnHTML(".s-main-slot .s-result-item", func(e *colly.HTMLElement) {
		// Extract product title
		title := e.ChildText("h2 a span")
		// Extract product price
		priceWhole := e.ChildText(".a-price .a-price-whole")
		priceFraction := e.ChildText(".a-price .a-price-fraction")
		price := priceWhole + "." + priceFraction

		if title != "" && price != "" {
			fmt.Printf("Product: %s | Price: $%s\n", title, price)
		}
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with error: %v\n", r.Request.URL, err)
	})

	// Start scraping
	err := c.Visit("https://www.amazon.com/s?k=laptops")
	if err != nil {
		log.Fatal("Failed to start scraping:", err)
	}

	// Wait for asynchronous tasks to complete
	c.Wait()
}

