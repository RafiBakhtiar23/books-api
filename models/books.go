package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
	Tahun    int    `json:"tahun"`
}
