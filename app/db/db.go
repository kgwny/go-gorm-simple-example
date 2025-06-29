package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// debugログ
	fmt.Println("[debug] dbUser = " + dbUser)
	fmt.Println("[debug] dbPassword = " + dbPassword)
	fmt.Println("[debug] dbName = " + dbName)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbUser,
		dbPassword,
		dbName,
	)

	// debugログ
	fmt.Println("[debug] dsn = " + dsn)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Retrying DB connection... (%d/10)\n", i)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		panic("Could not connect to database.")
	}

}
