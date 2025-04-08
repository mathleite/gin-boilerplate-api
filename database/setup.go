package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

func Setup() {
	db, err := gorm.Open(postgres.Open(getDatabaseDns()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("Database connection fails:", err)
	}
	Db = db
	log.Println("Database connection successfully.")
	migrateDatabase()
}

func getDatabaseDns() string {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, dbname, port)
}

func migrateDatabase() {
	if err := Db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Fail to migrate database: ", err)
	}
}
