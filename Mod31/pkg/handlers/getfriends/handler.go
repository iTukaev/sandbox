package getfriends

import (
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Input struct {
	TargetID string `json:"target_id"`
}

func NewHandler(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handle := &Handle{
		groupService: service,
	}
	return handle.GetFriends
}
type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	GetFriends(ID string) ([]string, error)
}

func (h *Handle) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

	userID := chi.URLParam(r, "userId")

	friends, err := h.groupService.GetFriends(userID)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user not found"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Friends of user %s - %v", userID, friends)))
}
