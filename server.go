package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"sahabatrental.com/stock_movement/v2/modules/stock_movement"
)

func main() {
	viper.SetConfigName("env")
	viper.SetConfigType("dotenv")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("where is the config file huh? pass it! %w", err))
	}

	println(viper.GetString("APP_KEY"))

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.POST("/create-movement", stock_movement.Create)

	r.Run()
}
