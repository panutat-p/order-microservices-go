package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/panutat-p/order-microservices-go/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const Port = ":8080"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", HealthCheckHandler)
	e.GET("/hello", HelloHandler)

	fmt.Println("http://localhost:8080/swagger/index.html")

	e.Logger.Fatal(e.Start(Port))
}

// HealthCheckHandler
// @Summary      Readiness Probe
// @Description  health check
// @Tags         health
// @Produce      plain
// @Success      200              {string}		string		"OK"
// @Router       /					[get]
func HealthCheckHandler(c echo.Context) error {
	fmt.Println("🟦 OK")
	return c.String(http.StatusOK, "OK")
}

// HelloHandler
// @Summary      Get hello world message
// @Description  hello
// @Tags         hello
// @Produce      json
// @Success      200  						{object}		HelloResponse
// @Router       /hello 			[get]
func HelloHandler(c echo.Context) error {
	fmt.Println("🟦 hello")
	return c.JSON(http.StatusOK, HelloResponse{Msg: "Hello, World!"})
}

type HelloResponse struct {
	Msg string `json:"msg"`
}
