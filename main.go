package main

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/controllers"
	"go-url-shortener/initializers"
	"go-url-shortener/migrations"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToPostgres()
	migrations.Migrate()
}

func main() {
	engine := gin.Default()

	controllers.ShortURLRoutes(engine)

	err := engine.Run()
	if err != nil {
		return
	}
}
