package ageupdate

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
)

type Input struct {
	NewAge int `json:"new_age"`
}

func NewHandle(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handle := &Handle{
		groupService: service,
	}
	return handle.AgeUpdate
}
type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	AgeUpdate(ID string, age int) (string, error)
}

func (h *Handle) AgeUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

	userID := chi.URLParam(r, "userId")

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	f := &Input{}
	if err := json.Unmarshal(content, f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := h.groupService.AgeUpdate(userID, f.NewAge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}