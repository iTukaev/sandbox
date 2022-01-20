package dbService

import (
	"bytes"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func (s *Service) GetAll() ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	documents, err := s.coll.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	u := new(User)
	for documents.Next(ctx) {
		if err = documents.Decode(u); err != nil {
			log.Printf("document %d decoding error %v", documents.ID(), err)
		}
		res, err := json.Marshal(u)
		if err != nil {
			log.Printf("document %d marshalling error %v", documents.ID(), err)
		}
		buffer.Write(res)
	}
	return buffer.Bytes(), nil
}
