package models

import "time"

type CarColor struct {
	Id        uint
	Color     string
	Title     string
	Stock     uint
	RealStock uint
	CarId     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
