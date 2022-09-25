package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/panutat-p/order-microservices-go/handler"
)

const Port = ":8080"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h, err := handler.New(e, Port)

	if err != nil {
		panic(err)
	}

	err = h.RunServer()
	if err != nil {
		panic(err)
	}
}
