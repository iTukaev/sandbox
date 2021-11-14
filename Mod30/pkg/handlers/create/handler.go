package create

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func (h *Handle) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect method"))
		return
	}
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

	inputPayload := &Input{}
	if err := json.Unmarshal(content, inputPayload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := h.groupService.CreateUser(inputPayload.Name, inputPayload.Age)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}
