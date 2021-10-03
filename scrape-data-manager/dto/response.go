package dto

type ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
