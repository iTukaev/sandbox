package getfriends

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type Input struct {
	TargetID int `json:"target_id"`
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
	GetFriends(ID int) ([]string, error)
}

func InternalError(w http.ResponseWriter, errStr string)  {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errStr))
}


func (h *Handle) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

	param := chi.URLParam(r, "userId")
	userID, err := strconv.Atoi(param)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	result, err := h.groupService.GetFriends(userID)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Friends of user ID: %d - %v", userID, result)))
}
