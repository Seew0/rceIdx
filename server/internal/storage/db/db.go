package db

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/seew0/rceIdx/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	client  *mongo.Client
	dbName  string
	timeout time.Duration
}

func NewStorage(client *mongo.Client, dbName string, timeout time.Duration) *Storage {
	return &Storage{
		client:  client,
		dbName:  dbName,
		timeout: timeout,
	}
}

func NewMongoClient(dbUri string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to MongoDB")
	}

	// Ping the primary
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "failed to ping MongoDB")
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func (s *Storage) GetUser(username string) (string, error) {
	return "", nil
}

func (s *Storage) InsertUser(user models.User) error {
	fmt.Println("sdahnjbdnsabdhasvbghdasbvgdhaebvghjadhjcbdvhcaghksbcvgkdfsabcghkfasbhkcbfsadhjklcahijldfb")
	_, err := s.client.Database(s.dbName).Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
