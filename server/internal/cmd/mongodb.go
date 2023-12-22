package main

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	db *mongo.Database
}
