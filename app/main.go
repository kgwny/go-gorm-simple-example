package main

import (
	"fmt"
	"my-app/db"
	"my-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("start.")
	db.Init()

	server := echo.New()
	routes.RegisterRoutes(server)

	server.Logger.Fatal(server.Start(":8080"))
	fmt.Println("end.")
}
