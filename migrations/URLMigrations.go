package migrations

import (
	"go-url-shortener/initializers"
	"go-url-shortener/models"
	"log"
)

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.Urls{})
	if err != nil {
		log.Fatal("Something wrong when migration data to the database")
		return
	}
	log.Println("Migrations ran with no problems")
}
