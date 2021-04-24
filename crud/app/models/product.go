package models

import (
	"gorm.io/gorm"
)

// Product struct to describe product object.
type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Type  string `json:"type"`
	Prize int    `json:"prize"`
}
