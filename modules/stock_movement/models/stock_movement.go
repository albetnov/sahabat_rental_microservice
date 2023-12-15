package models

import (
	"gorm.io/gorm"
	"time"
)

type StockMovement struct {
	gorm.Model
	CarColorId  uint
	SourceStock int
	ToStock     int
	Desc        string
	MovedAt     time.Time
}
