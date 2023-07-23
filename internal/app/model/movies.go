package model

import "gorm.io/gorm"

type Movies struct {
	ID          int     `gorm:"primaryKey"`
	Title       string  `gorm:"column:title;not null;size:256;index;unique"`
	Description string  `gorm:"column:description;type:text;not null"`
	Rating      float64 `gorm:"column:rating;not null"`
	Image       string  `gorm:"column:image;not null"`
	gorm.Model
}
