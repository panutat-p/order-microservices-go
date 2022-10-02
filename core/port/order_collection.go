package port

import "github.com/panutat-p/order-microservices-go/core/domain"

type OrderStore interface {
	Find(itemType string) ([]domain.Order, error)
	Save(order *domain.Order) error
}
