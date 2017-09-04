package models

import (
	"time"
)

type Media struct {
	ID        uint `gorm:"primary_key" json:"id"`
	RealPath  string `json:"realPath"`
	Path 	  string `json:"path"`
	CreatedAt time.Time `json:"createdAt"`
}