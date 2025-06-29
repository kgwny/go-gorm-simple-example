package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo) {

	// http://localhost:8080/user
	server.GET("/user", getAllUsers)

	// http://localhost:8080/user/{userId}
	server.GET("/user/:userId", getUserById)
}
