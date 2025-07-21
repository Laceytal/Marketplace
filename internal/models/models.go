package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string
}

type Ad struct {
	gorm.Model
	Title       string
	Description string
	ImageURL    string
	Price       float64
	UserID      uint
}
