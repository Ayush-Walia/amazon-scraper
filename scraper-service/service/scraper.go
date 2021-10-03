package service

import (
	"log"
	"strconv"
	"strings"

	"github.com/Ayush-Walia/amazon-scraper/scraper-service/dto"
	"github.com/Ayush-Walia/amazon-scraper/scraper-service/utils"
	"github.com/gocolly/colly"
)

// Scrapper service scrapes the web page from the given url.
func Scrapper(URL string) (dto.ResponseMessage, dto.ScrapedData) {
	var scrapingResponse dto.ResponseMessage
	var scrapedData dto.ScrapedData
	var productData dto.ProductDetails
	collyCollector := colly.NewCollector()

	collyCollector.OnRequest(func(r *colly.Request) {
		log.Println("Scrapping ", r.URL)
	})

	collyCollector.OnHTML("#filter-info-section > div", func(c *colly.HTMLElement) {
		// Scrape Product reviews count.
		reviews := c.ChildText("#filter-info-section > div > span")
		review := strings.Split(reviews, " ")
		productData.TotalReviews, _ = strconv.Atoi(strings.Replace(review[4], ",", "", -1))
	})

	collyCollector.OnHTML("div#dp", func(c *colly.HTMLElement) {
		// Scrape Product Name, Product Image URL and Product Price.
		productData.Name = c.ChildText("#productTitle")
		productData.ImageURL = c.ChildAttr("#landingImage", "src")
		productData.Price = c.ChildText("#priceblock_ourprice")
		if productData.Price == "" {
			productData.Price = "Price Unavailable. Product is out of stock or Product has multiple choices!"
		}

		// Scrape Product Description.
		var descriptionArray []string
		c.ForEach("div#feature-bullets", func(_ int, cHTMLElement *colly.HTMLElement) {
			descriptionArray = append(descriptionArray, cHTMLElement.ChildText("span.a-list-item"))
		})
		for _, description := range descriptionArray {
			productData.Description = strings.Split(description, "\n\n\n")
		}

		reviewPage := c.ChildAttr("#cr-pagination-footer-0 > a", "href")
		reviewPage = c.Request.AbsoluteURL(reviewPage)
		err := collyCollector.Visit(reviewPage)
		utils.CheckError(err)
	})

	err := collyCollector.Visit(URL)

	if err != nil {
		scrapingResponse.Success = false
		scrapingResponse.Message = err.Error()
		return scrapingResponse, scrapedData
	}

	scrapedData.URL = URL
	scrapedData.Product = productData
	scrapingResponse.Success = true
	scrapingResponse.Message = "Successfully scraped the page."
	return scrapingResponse, scrapedData
}
