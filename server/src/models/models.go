package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	Title string `gorm:"text;not null;default:null" json:"title"`
	Image string `gorm:"text;not null;default:null" json:"image"`
}
