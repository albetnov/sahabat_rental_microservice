package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Gorm *gorm.DB
}

var Database = DB{}
var Gorm *gorm.DB

func (db *DB) Configure() {
	viper.SetDefault("DB_HOST", "127.0.0.1")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_CHARSET", "latin1")
}

func (db *DB) Connect(
	username string,
	password string,
	host string,
	port string,
	name string,
	charset string,
) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
		charset,
	)

	var err error

	db.Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Gorm = db.Gorm

	if err != nil {
		panic("Failed to connect to database!")
	}
}
