package handlers

import (
	"log"
	"net/http"

	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/dto"
	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/service"
	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/utils"
)

// HandleSaveProduct handles saving product in database.
func HandleSaveProduct(w http.ResponseWriter, r *http.Request) {
	var saveProductRequest dto.ScrapedData
	utils.GetStructFromRequest(w, r, &saveProductRequest)
	response := service.SaveData(saveProductRequest)

	if !response.Success {
		log.Println("Failed to write data to DB: ", response.Message)
		utils.RespondWithError(w, 400, response.Message)
		return
	}
	utils.RespondWithJSON(w, 200, response)
}
