package entity

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	Collections *mongo.Collection
}