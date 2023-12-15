package models

import (
	"time"
)

type StockRestoration struct {
	ID                 uint `gorm:"primaryKey"`
	StockMovementId    uint
	StockMovement      StockMovement
	NewStockMovementId uint
	NewStockMovement   StockMovement
	ScheduledAt        time.Time
	RestoredAt         *time.Time
	Status             string
	SourceModule       string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
