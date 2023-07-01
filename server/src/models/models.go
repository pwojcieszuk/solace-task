package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	Id    int    `gorm:"primaryKey" json:"id"`
	Title string `gorm:"text;not null;default:null" json:"title"`
	Image string `gorm:"text;not null;default:null" json:"image"`
}
