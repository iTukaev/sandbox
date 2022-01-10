package dbService

import (
	"context"
	"log"
)

func Stop(ctx context.Context) error {
	<-ctx.Done()
	if err = Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	log.Println("MongoDB stopped")
	return nil
}
