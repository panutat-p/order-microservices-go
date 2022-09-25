package port

import "github.com/panutat-p/order-microservices-go/core/order"

type OrderStore interface {
	Save(order *order.Order) error
}
