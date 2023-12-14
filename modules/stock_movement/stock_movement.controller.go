package stock_movement

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"sahabatrental.com/stock_movement/v2/modules/stock_movement/models"
	"sahabatrental.com/stock_movement/v2/utils"
)

func Create(c *gin.Context) {
	var json StockMovement

	if err := c.ShouldBindJSON(&json); err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			utils.Response.Validation(c, &ve)

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload Invalid",
			"error":   err.Error(),
		})

		return
	}

	earning := models.Earning{}

	if result := utils.Gorm.Preload("CarColor").Where("id = ?", json.EarningId).First(&earning); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Earning ID not found!",
			})
			return
		}
	}

	go func() {
		err := utils.Gorm.Transaction(func(tx *gorm.DB) error {
			stockMovement := models.StockMovement{
				CarColorId:  earning.CarColorId,
				SourceStock: earning.CarColor.RealStock,
				ToStock:     earning.CarColor.RealStock - earning.Qty,
				Desc:        json.Desc,
			}

			if result := utils.Gorm.Create(&stockMovement); result.Error != nil {
				return result.Error
			}

			stockRestoration := models.StockRestoration{
				StockMovementId: stockMovement.ID,
				ScheduledAt:     json.ReturnedAt,
				Status:          "pending",
				SourceModule:    "all.navigation.earning",
			}

			if result := utils.Gorm.Create(&stockRestoration); result.Error != nil {
				return result.Error
			}

			return nil
		})

		if err != nil {
			// fire webhook
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock movement Processed!",
	})
}
