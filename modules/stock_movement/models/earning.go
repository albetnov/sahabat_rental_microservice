package models

import "time"

type Earning struct {
	ID              uint `gorm:"primaryKey"`
	Code            string
	CarColorsId     uint
	CarColor        CarColor `gorm:"foreignKey:CarColorsId"`
	CustomerId      uint
	StockMovementId uint
	StockMovement   StockMovement
	Qty             uint
	Total           float64
	Note            string
	RentAt          time.Time
	ReturnedAt      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
