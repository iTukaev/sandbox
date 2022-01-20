package create

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Input struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewHandler(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handler := &Handle {
		groupService: service,
	}
	return handler.Create
}

type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	CreateUser(name string, age int) (string, error)
}

func InternalError(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errStr))
}


func (h *Handle) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}

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

	ID, err := h.groupService.CreateUser(inputPayload.Name, inputPayload.Age)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User name: %s created with ID: %s", inputPayload.Name, ID)))
}
