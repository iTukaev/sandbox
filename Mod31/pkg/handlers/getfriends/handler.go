package getfriends

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Input struct {
	TargetID string `json:"target_id"`
}

func NewHandle(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handle := &Handle{
		groupService: service,
	}
	return handle.GetFriends
}
type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	GetFriends(ID string) (string, error)
}

func (h *Handle) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

	userID := chi.URLParam(r, "userId")

	result, err := h.groupService.GetFriends(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
