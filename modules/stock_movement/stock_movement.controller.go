package stock_movement

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock movement created!",
	})
}
