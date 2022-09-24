package main

import (
	"order-microservices-go/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const Port = ":8080"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.HelloHandler)

	e.Logger.Fatal(e.Start(Port))
}
