package main

import (
	"github.com/labstack/echo"
	"goAPI/routing"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
  // e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(bodyDumpHandler))

	// routing
	e.POST("/user",routing.BaseAPI_user())

	e.Start(":8080")
}
