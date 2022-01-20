package dbService

import (
	"bytes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (s *Service) GetAll() ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	documents, err := s.coll.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var builder bytes.Buffer
	for documents.Next(ctx) {
		builder.WriteString(documents.Current.String())
	}
	return builder.Bytes(), nil
}
