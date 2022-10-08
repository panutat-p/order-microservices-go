package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/panutat-p/order-microservices-go/core/domain"
	"github.com/panutat-p/order-microservices-go/core/mocks"
)

func TestOrderService_List(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mocks.NewMockOrderStorer(ctrl)

	service := NewOrderService(m)

	given := "banana"
	want := []domain.Order{{Country: "Thailand"}}

	m.EXPECT().Find(given).Return(want, nil).Times(1)

	got, err := service.List(given)

	if err != nil {
		t.Error(err)
	}

	if len(got) != 1 {
		t.Error("want", len(want), "but got", len(got))
	}
}
