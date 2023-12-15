package stock_movement

type StockMovement struct {
	EarningId uint   `json:"earning_id" binding:"required,number"`
	Desc      string `json:"desc" binding:"required,max=225"`
}
