package models

import "gorm.io/gorm"

type Tile struct {
	gorm.Model
	Title        string   `json:"title"`
	Directors    []string `json:"directors"`
	Actors       []string `json:"actors"`
	Descriptions string   `json:"description"`
}
