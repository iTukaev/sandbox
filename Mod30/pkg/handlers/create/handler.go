package create

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Input struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Handle struct {
	GroupService groupInterface
}

type groupInterface interface {
	CreateUser(name string, age int) (string, error)
	MakeFriend(TargetID int, SourceID int) (string, error)
}

func (h *Handle) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	inputPayload := &Input{}
	if err := json.Unmarshal(content, inputPayload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := h.GroupService.CreateUser(inputPayload.Name, inputPayload.Age)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}
