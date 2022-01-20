package makefriend

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type Input struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
}

func NewHandler(service groupInterface) func(w http.ResponseWriter, r *http.Request) {
	handler := &Handle {
		groupService: service,
	}
	return handler.MakeFriend
}

type Handle struct {
	groupService groupInterface
}

type groupInterface interface {
	MakeFriend(TargetID string, SourceID string) error
}

func InternalError(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errStr))
}

func (h *Handle) MakeFriend(w http.ResponseWriter, r *http.Request) {
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

	f := &Input{}
	if err := json.Unmarshal(content, f); err != nil {
		InternalError(w, err.Error())
		return
	}

	err = h.groupService.MakeFriend(f.TargetID, f.SourceID)
	if errors.Is(err, mongo.ErrNoDocuments) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("user not found, %v", err)))
		return
	}
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
