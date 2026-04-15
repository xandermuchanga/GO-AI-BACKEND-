package models

type MessageRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type MessageResponse struct {
	Response string `json:"response"`
	Category string `json:"category"`
}

type AIResponse struct {
	Response string `json:"response"`
	Category string `json:"category"`
}
