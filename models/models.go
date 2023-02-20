package models

import "gorm.io/gorm"

type BodyJsonRequest struct {
	Data string `json:"data"`
}

type SaveObject struct {
	gorm.Model
	Input  string `json:"Input"`
	Output string `json:"Output"`
}
