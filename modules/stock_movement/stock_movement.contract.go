package stock_movement

import "time"

type StockMovement struct {
	EarningId  uint      `json:"earning_id" binding:"required,number"`
	ReturnedAt time.Time `json:"returned_at" binding:"required,gt"`
	Module     string    `json:"module" binding:"required,oneof=earnings"`
	Desc       string    `json:"desc" binding:"required,max=225"`
}
