package create

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
	CreateUser(name string, age int) int
}

func InternalError(w http.ResponseWriter, errStr string)  {
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

	input := &Input{}
	if err = json.Unmarshal(content, input); err != nil {
		InternalError(w, err.Error())
		return
	}

	ID := h.groupService.CreateUser(input.Name, input.Age)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(ID)))
}