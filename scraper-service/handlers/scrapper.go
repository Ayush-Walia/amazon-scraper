package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ayush-Walia/amazon-scraper/scraper-service/dto"
	"github.com/Ayush-Walia/amazon-scraper/scraper-service/service"
	"github.com/Ayush-Walia/amazon-scraper/scraper-service/utils"
)

// HandlePageScraping handles the Page Scraping. It reads url from request payload and calls scrape service.
// After scraping the info from the page, it calls another API to save the data.
func HandlePageScraping(w http.ResponseWriter, r *http.Request) {
	var scrapeRequest dto.ScrapePageRequest

	// Convert request to go struct.
	utils.GetStructFromRequest(w, r, &scrapeRequest)

	if len(scrapeRequest.URL) == 0 {
		utils.RespondWithError(w, 400, "empty URL")
		return
	}

	// Call scrapper service, respond back with error in case of error else respond back with scrapped data.
	status, data := service.Scrapper(scrapeRequest.URL)
	if !status.Success {
		utils.RespondWithError(w, 400, "The scraping failed with error "+status.Message)
		return
	}

	// Calling scrape-data-service API for saving the data in DB.
	apiPayload, _ := json.Marshal(data)
	scrapeDataManagerAPI := "http://scrape-data-manager:8080/save_product"
	log.Println("Making API call to save the scraped data")
	_, err := http.Post(scrapeDataManagerAPI, "application/json", bytes.NewBuffer(apiPayload))
	if err != nil {
		log.Println("The HTTP request failed with error ", err)
		utils.RespondWithError(w, 500, "Error saving data in DB")
		return
	}
	log.Println("Successfully scraped and saved data in DB")
	utils.RespondWithJSON(w, 200, data)
}
