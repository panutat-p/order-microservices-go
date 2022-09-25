package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/panutat-p/order-microservices-go/core/order"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderStore struct {
	*mongo.Collection
}

const MongoTimeout = 10

func NewOrderStore(uri string) *OrderStore {
	if uri == "" {
		log.Fatal("🟥 Failed to load ENV file")
	}

	opt := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		defer log.Fatal(err)
	}

	if err != nil {
		fmt.Println("🟥 Failed to connect to MongoDB")
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("order-microservices-go").Collection("orders")
	return &OrderStore{Collection: coll}
}

func (s *OrderStore) Save(order *order.Order) error {
	_, err := s.Collection.InsertOne(context.Background(), order)
	return err
}
