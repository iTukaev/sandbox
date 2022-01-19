package ageupdate

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
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
	AgeUpdate(ID int, age int) (string, error)
}

func InternalError(w http.ResponseWriter, errStr string)  {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errStr))
}


func (h *Handle) AgeUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
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

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		InternalError(w, err.Error())
		return
	}
	r.Body.Close()

	f := &Input{}
	if err := json.Unmarshal(content, f); err != nil {
		InternalError(w, err.Error())
		return
	}

	name, err := h.groupService.AgeUpdate(userID, f.NewAge)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("User ID: %d, Name: %s, age chang successfull", userID, name )))
}