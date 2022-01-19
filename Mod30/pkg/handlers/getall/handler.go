package getall

import (
	"net/http"
)

func NewHandle(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handle := &Handle{
		groupService: service,
	}
	return handle.GetAll
}

type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	GetAll() ([]byte, error)
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

	body, err := h.groupService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
