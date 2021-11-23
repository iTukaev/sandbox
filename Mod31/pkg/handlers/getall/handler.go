package getall

import (
	"bytes"
	"net/http"
)

func NewHandler(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handle := &Handle{
		groupService: service,
	}
	return handle.GetAll
}

type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	GetAll() (*bytes.Buffer, error)
}

func (h *Handle) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

	buf, err := h.groupService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}
