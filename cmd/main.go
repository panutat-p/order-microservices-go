package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/panutat-p/order-microservices-go/core/service"
	"github.com/panutat-p/order-microservices-go/handler"
	"github.com/panutat-p/order-microservices-go/store/mongodb"
)

const Port = ":8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("🟥 ENV, err:", err)
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	orderStore := mongodb.NewOrderStore(os.Getenv("MONGODB_URI"))
	orderService := service.NewOrderService(orderStore)
	h, err := handler.New(e, Port, orderService)

	if err != nil {
		panic(err)
	}

	err = h.RunServer()
	if err != nil {
		panic(err)
	}
}
