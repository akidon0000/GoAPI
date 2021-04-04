package main

import (
	"github.com/labstack/echo"
	"goAPI/routing"
)

func main() {
	e := echo.New()

	// routing
	e.POST("/user",routing.BaseAPI_user())

	e.Start(":8080")
}
