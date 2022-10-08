package port

import "github.com/panutat-p/order-microservices-go/core/domain"

//go:generate mockgen -destination=../mocks/order_storer.go -package=mocks github.com/panutat-p/order-microservices-go/core/port OrderStorer
type OrderStorer interface {
	Find(itemType string) ([]domain.Order, error)
	Save(order *domain.Order) error
}
