package model

type ShortenUrlRequestModel struct {
	Url string `json:"url" binding:"required"`
}
