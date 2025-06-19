package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDB() (err error) {
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8m64&parseTime=True&loc=local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(mysql.Open(connectInfo), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Database connected.")
	return nil
}
