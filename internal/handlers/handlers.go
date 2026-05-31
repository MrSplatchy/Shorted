package handlers

import (
	"errors"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/models"

	"github.com/asaskevich/govalidator"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func CreateUrl(c *gin.Context) {
	var req models.UrlRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok := govalidator.IsURL(req.Url)
	if ok != true {
		c.JSON(400, gin.H{"error": "This is not an URL, please retry"})
		return
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

	// Incase if the record already exists
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, res)
		return
	}

	err = config.DB.Create(&shape).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		shape.ShortCode = uniuri.NewLen(6)
	}

	c.JSON(http.StatusCreated, res)
}

func RetrieveUrl(c *gin.Context) {
	var req models.UrlRequest

	req.Shortcode = c.Param("shortcode")

	shape := models.Urls{
		ShortCode: req.Shortcode,
	}

	err := config.DB.Where("short_code = ?", req.Shortcode).First(&shape).Error

	res := models.UrlResponse{
		ID:        shape.ID,
		Url:       shape.Url,
		ShortCode: shape.ShortCode,
		CreatedAt: shape.CreatedAt,
		UpdatedAt: shape.UpdatedAt,
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusBadRequest, "error")
		return
	}

	c.JSON(http.StatusOK, res)

}
