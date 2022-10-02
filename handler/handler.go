package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/panutat-p/order-microservices-go/core/service"
	_ "github.com/panutat-p/order-microservices-go/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const ReadTimeOut = 3

var handlerErr = errors.New("cannot init echo web server")

type Handler struct {
	Router       *echo.Echo
	Port         string
	OrderService service.OrderService
}

func New(e *echo.Echo, port string, orderService service.OrderService) (*Handler, error) {
	if e == nil {
		return nil, handlerErr
	}

	return &Handler{
		Router:       e,
		Port:         port,
		OrderService: orderService,
	}, nil
}

// RunServer
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
func (h *Handler) RunServer() error {
	h.Router.GET("/swagger/*", echoSwagger.WrapHandler)
	h.Router.GET("/", h.healthCheck)
	h.Router.POST("/reflect", h.reflectReq)
	h.Router.GET("/list", h.listOrders)
	h.Router.POST("/create", h.createOrder)

	s := http.Server{
		ReadTimeout: ReadTimeOut * time.Second,
		Addr:        h.Port,
		Handler:     h.Router,
	}

	fmt.Println("http://localhost:8080/swagger/index.html")

	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
