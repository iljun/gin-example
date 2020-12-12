package global

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)
const (
	DB_USERNAME = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME = "DB_NAME"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitializeDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.New(extractorDataBaseInfo()), &gorm.Config{})

	if err != nil {
		fmt.Println("db err: ", err)
		panic(err)
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func extractorDataBaseInfo() postgres.Config {
	dbUser := getEnv(DB_USERNAME)
	dbPassword := getEnv(DB_PASSWORD)
	dbSchema := getEnv(DB_NAME)

	info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbSchema)

	return postgres.Config{
		DSN: info,
		PreferSimpleProtocol: true,
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		message := fmt.Sprintf("OS env not found key : %s", key)
		panic(message)
	}
	return value
}
