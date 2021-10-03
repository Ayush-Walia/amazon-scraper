package handlers

import (
	"net/http"

	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/dto"
	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/utils"
)

// HandleHealthCheck handles health check request.
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	var response dto.ResponseMessage
	response.Success = true
	response.Message = "health check is awesome!"
	utils.RespondWithJSON(w, http.StatusOK, response)
}
