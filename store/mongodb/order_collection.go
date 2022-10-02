package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/panutat-p/order-microservices-go/core/domain"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderStore struct {
	coll *mongo.Collection
}

const MongoTimeout = 10
const DatabaseName = "order_microservices"
const CollectionName = "orders"

func NewOrderStore(uri string) OrderStore {
	if uri == "" {
		log.Fatal("🟥 Failed to load ENV file")
	}

	opt := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), MongoTimeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println("🟥 Failed to connect to MongoDB")
		panic(err)
	}

	coll := client.Database(DatabaseName).Collection(CollectionName)
	return OrderStore{coll}
}

func (s OrderStore) Find(itemType string) ([]domain.Order, error) {
	filter := bson.D{primitive.E{Key: "item_type", Value: itemType}}
	var d bson.D
	err := s.coll.FindOne(context.Background(), filter).Decode(&d)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("🟥 Find() did not match any documents, err:", err)
			return nil, err
		}
		fmt.Println("🟥 Find(), err:", err)
		return nil, err
	}

	var orders []domain.Order
	bsonBytes, _ := bson.Marshal(d)
	err = bson.Unmarshal(bsonBytes, orders)
	if err != nil {
		fmt.Println("🟥 Unmarshal(), err:", err)
		return nil, err
	}
	return orders, nil
}

func (s OrderStore) Save(order *domain.Order) error {
	_, err := s.coll.InsertOne(context.Background(), order)
	if err != nil {
		fmt.Println("🟥 Save() err:", err)
		return err
	}
	return nil
}
