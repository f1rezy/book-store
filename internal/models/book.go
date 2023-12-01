package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title      string `json:"title"`
	Price      uint   `json:"price"`
	Annotation string `json:"annotation"`
	Category   string `json:"category"`
}
