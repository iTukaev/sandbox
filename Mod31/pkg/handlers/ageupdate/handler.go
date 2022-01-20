package ageupdate

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

type Input struct {
	NewAge int `json:"new_age"`
}

func NewHandler(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handle := &Handle{
		groupService: service,
	}
	return handle.AgeUpdate
}
type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	AgeUpdate(ID string, age int) error
}

func InternalError(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errStr))
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
		InternalError(w, err.Error())
		return
	}
	defer r.Body.Close()

	inputPayload := &Input{}
	if err := json.Unmarshal(content, inputPayload); err != nil {
		InternalError(w, err.Error())
		return
	}

	err = h.groupService.AgeUpdate(userID, inputPayload.NewAge)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}