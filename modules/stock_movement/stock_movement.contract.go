package stock_movement

type StockMovement struct {
	EarningId uint   `json:"earning_id" binding:"required,number"`
	Module    string `json:"module" binding:"required,oneof=earnings"`
	Desc      string `json:"desc" binding:"required,max=225"`
}
