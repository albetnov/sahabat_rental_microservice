package stock_movement

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var json StockMovement

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock movement created!",
	})
}
