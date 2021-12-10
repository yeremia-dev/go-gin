package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/yeremia-dev/go-gin/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is used to creating database connection
func SetupDatabaseConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	//This line code is used to generate table on the database automatically.
	//If the entity will be generate has relation to another entity, the required entity will be generated too
	db.AutoMigrate(&entity.Book{}, &entity.User{})

	return db
}

//CloseDatabaseConnection is used to close connection to database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}
	dbSql.Close()
}
