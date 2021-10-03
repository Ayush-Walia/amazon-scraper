package dto

type ProductDetails struct {
	Name         string   `json:"name"`
	ImageURL     string   `json:"image_url,omitempty"`
	Description  []string `json:"description,omitempty"`
	Price        string   `json:"price,omitempty"`
	TotalReviews int      `json:"total_reviews,omitempty"`
}

type ScrapedData struct {
	URL     string         `json:"url"`
	Product ProductDetails `json:"product"`
}
