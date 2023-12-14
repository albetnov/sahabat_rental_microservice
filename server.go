package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sahabatrental.com/stock_movement/v2/modules/stock_movement"
	"sahabatrental.com/stock_movement/v2/utils"
	"time"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	viper.SetDefault("APP_PORT", "3000")
	viper.SetDefault("TRUSTED_PROXIES", []string{"127.0.0.1"})

	// configure viper default env settings
	utils.Database.Configure()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("where is the config file huh? pass it! %w", err))
	}

	if viper.GetString("APP_KEY") == "" {
		panic("APP_KEY must be passed!")
	}

	// connect to database
	utils.Database.Connect(
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_CHARSET"),
	)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if err := r.SetTrustedProxies(viper.GetStringSlice("TRUSTED_PROXIES")); err != nil {
		log.Fatalf("Trusted Proxes is invalid: %s", err.Error())
	}

	r.GET("/health", func(c *gin.Context) {
		result, err := time.Parse("02-01-2006", "01-01-2023")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"time":    result,
		})
	})

	r.Use(func(c *gin.Context) {
		var key string

		key = c.GetHeader("Authorization")

		if key == "" {
			key = c.Query("app_key")
		}

		if key == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "APP_KEY is missing",
				"hint":    "Either put key in authorization header or use app_key query string",
			})

			c.Abort()
			return
		}

		if s, err := base64.StdEncoding.DecodeString(key); err == nil && string(s) == viper.GetString("APP_KEY") {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid APP_KEY",
		})

		c.Abort()
	})

	r.POST("/create-movement", stock_movement.Create)

	if err := r.Run(fmt.Sprintf(":%s", viper.GetString("APP_PORT"))); err != nil {
		fmt.Println(err.Error())
		panic("okay the port may be blocked? you can add APP_PORT env for alternate port though")
	}
}
