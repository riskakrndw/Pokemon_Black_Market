package config

import (
	"os"
	"project/pbm/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	var err error
	connectionString := "root:welcome12345@tcp(localhost:3306)/blackmarket?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.DetailAbility{})
	DB.AutoMigrate(&models.DetailTransaction{})
	DB.AutoMigrate(&models.DetailType{})
	DB.AutoMigrate(&models.Level{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.Pokemon{})
	DB.AutoMigrate(&models.PokemonAbility{})
	DB.AutoMigrate(&models.PokemonType{})
	DB.AutoMigrate(&models.User{})
}

func ConfigTest() (*gorm.DB, error) {
	var err error
	connectionStringTest := "root:welcome12345@tcp(localhost:3306)/blackmarket_testing?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionStringTest), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}
