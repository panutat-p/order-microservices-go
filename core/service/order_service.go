package service

import (
	"github.com/panutat-p/order-microservices-go/core/order"
	"github.com/panutat-p/order-microservices-go/core/port"
)

type OrderService struct {
	store port.OrderStore
}

func (s *OrderService) Create(order *order.Order) error {
	err := s.store.Save(order)
	if err != nil {
		return err
	}
	return nil
}
