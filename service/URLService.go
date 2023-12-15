package service

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go-url-shortener/dto"
	"go-url-shortener/initializers"
	"go-url-shortener/models"
	"net/http"
	"net/url"
)

func GenerateShortURL(c *gin.Context) {
	var urlRequest dto.GenerateShortURLRequest

	err := c.BindJSON(&urlRequest)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error while mapping body into the DTO"})
		return
	}

	// Parse the long URL
	parsedURL, err := url.Parse(urlRequest.LongURL)
	if err != nil {
		// Handle invalid long URL
		c.IndentedJSON(http.StatusBadRequest, "Invalid long URL")
		return
	}

	// Validate the parsed URL
	isValid := parsedURL.IsAbs() && parsedURL.Host != "" && parsedURL.Scheme != ""
	if !isValid {
		// Handle invalid long URL
		c.IndentedJSON(http.StatusBadRequest, "Invalid long URL")
		return
	}

	var checkIfLongURLAlreadyExists models.Urls
	initializers.DB.Where("long_url = ?", urlRequest.LongURL).First(&checkIfLongURLAlreadyExists)
	if checkIfLongURLAlreadyExists.ID != 0 {
		c.IndentedJSON(http.StatusCreated, gin.H{"shortURL": checkIfLongURLAlreadyExists.ShortUrl})
		return
	}

	// Use a hashing algorithm to generate a unique short URL
	hash := sha1.Sum([]byte(urlRequest.LongURL))
	shortURL := base64.URLEncoding.EncodeToString([]byte(hash[:10]))

	newUrl := models.Urls{
		LongUrl:  urlRequest.LongURL,
		ShortUrl: shortURL,
	}

	initializers.DB.Create(&newUrl)

	c.IndentedJSON(http.StatusCreated, gin.H{"shortURL": newUrl.ShortUrl})
}

func GetLongUrlBasedOnSmall(c *gin.Context) {
	var urlByShortURL models.Urls

	shortUrl := c.Param("shortUrl")

	first := initializers.DB.Where("short_url = ?", shortUrl).First(&urlByShortURL)

	if first.Error != nil {
		if err := first.Error; err != pgx.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error fetching long URL"})
			return
		}

		// Handle the case where short URL doesn't exist in the database
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, urlByShortURL.LongUrl)
}
