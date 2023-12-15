package models

import "gorm.io/gorm"

type Urls struct {
	gorm.Model
	LongUrl  string
	ShortUrl string `gorm:"unique"`
}
