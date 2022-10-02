package service

import (
	"github.com/panutat-p/order-microservices-go/core/domain"
	"github.com/panutat-p/order-microservices-go/core/port"
)

type OrderService struct {
	store port.OrderStorer
}

func NewOrderService(orderStore port.OrderStorer) OrderService {
	return OrderService{
		store: orderStore,
	}
}

func (s *OrderService) List(itemType string) ([]domain.Order, error) {
	orders, err := s.store.Find(itemType)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) Create(order *domain.Order) error {
	err := s.store.Save(order)
	if err != nil {
		return err
	}
	return nil
}
