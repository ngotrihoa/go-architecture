package main

import (
	"log"
	"os"
	"todo-list/server"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_CONNECTION")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.Debug()
	log.Println("DB Connection Success:", db)

	s := server.NewServer(db)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
