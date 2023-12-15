package controllers

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/service"
)

func ShortURLRoutes(r *gin.Engine) {
	route := r.Group("/short-url")
	route.POST("", service.GenerateShortURL)
	route.GET("/:shortUrl", service.GetLongUrlBasedOnSmall)
}
