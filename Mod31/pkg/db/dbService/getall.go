package dbService

import (
	"bytes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) GetAll() (*bytes.Buffer, error) {
	var buf bytes.Buffer

	documents, err := s.coll.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	for documents.Next(context.TODO()) {
		buf.WriteString(documents.Current.String() + "\n")
	}
	return &buf, nil
}
