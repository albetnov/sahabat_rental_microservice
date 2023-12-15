package models

import "time"

type CarColor struct {
	ID        uint `gorm:"primaryKey"`
	Color     string
	Title     string
	Stock     uint
	RealStock int
	CarId     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
