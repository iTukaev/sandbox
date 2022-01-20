package deleteuser

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type Input struct {
	TargetID string `json:"target_id"`
}

func NewHandler(service groupInterface) func(w http.ResponseWriter, r *http.Request)  {
	handle := &Handle{
		groupService: service,
	}
	return handle.Delete
}

type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	DeleteUser(ID string) error
}

func InternalError(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errStr))
}


func (h *Handle) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	err = h.groupService.DeleteUser(inputPayload.TargetID)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("User not found"))
		return
	}
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User ID %s was deleted", inputPayload.TargetID)))
}
