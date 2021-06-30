package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type ConfigDatabase struct {
	AppName  string `env:"APP_NAME" envDefault:"app"`
	AppEnv   string `env:"APP_ENV" envDefault:"dev"`
	Port     string `env:"MY_APP_PORT" envDefault:"1323"`
	Host     string `env:"HOST" envDefault:"localhost"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

var cfg ConfigDatabase

var (
	db       *gorm.DB
	username = "docker"
	password = "docker"
	host     = "127.0.0.1:3306"
	schema   = "database"
)

func Connect() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, host, schema)

	var err error
	d, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	db = d

	log.Println("database successfully configured")
}

func GetDB() *gorm.DB {
	return db
}
