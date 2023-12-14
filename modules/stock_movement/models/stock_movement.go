package models

import (
	"gorm.io/gorm"
	"time"
)

type StockMovement struct {
	gorm.Model
	CarColorId  uint
	SourceStock uint
	ToStock     uint
	Desc        string
	MovedAt     time.Time
}
