package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"shortener/internal/config"
	"shortener/internal/models"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func CreateUrl(c *gin.Context) {
	var req models.UrlRequest

	req.Url = c.Query("url")
	_, err := url.Parse(req.Url)
	if err != nil {
		c.String(400, "Bad request")
	}

	shape := models.Urls{
		Url:       req.Url,
		ShortCode: uniuri.NewLen(6),
	}

	err = config.DB.Where("url = ?", req.Url).First(&shape).Error
	// Setup a response
	res := models.UrlResponse{
		ID:        shape.ID,
		Url:       shape.Url,
		ShortCode: shape.ShortCode,
		CreatedAt: shape.CreatedAt,
		UpdatedAt: shape.UpdatedAt,
	}
	// If It cannot find a Record, it will not populate the current shape, so we can add a new one
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = config.DB.Create(&shape).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			shape.ShortCode = uniuri.NewLen(7)
		}

		c.JSON(http.StatusCreated, res)
		return
	}
	//Normally shouldn't be executed if it is a new URL
	c.JSON(http.StatusCreated, res)
}
