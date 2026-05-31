package models

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	Url       string
	ShortCode string `gorm:"uniqueIndex"`
}

type UrlRequest struct {
	Url       string `json:"url" binding:"required,url"`
	Shortcode string
}

type UrlResponse struct {
	ID        uint      `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
