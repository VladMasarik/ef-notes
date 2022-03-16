package models

import (
	"gorm.io/gorm"
)

// Basic model of a note that is holding some text.
type Note struct {
	gorm.Model
	Text string `json:"text"`
}
