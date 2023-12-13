package stock_movement

import "time"

type StockMovement struct {
	EarningId  uint      `json:"earning_id" binding:"required" validate:"required,numeric"`
	ReturnedAt time.Time `json:"returned_at" binding:"required" validate:"required,datetime=2006-01-02,gt"`
	Module     string    `json:"module" binding:"required" validate:"required,oneof=earnings"`
}
