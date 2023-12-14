package models

import "time"

type Earning struct {
	Id              uint
	Code            string
	CarColorId      uint
	CarColor        CarColor
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
