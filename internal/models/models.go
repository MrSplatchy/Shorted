package models

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	Url       string `gorm:"not null"`
	ShortCode string `gorm:"uniqueIndex"`
}

type UrlRequest struct {
	Url string `json:"url" binding:"required,url"`
}

type UrlResponse struct {
	ID        uint      `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
