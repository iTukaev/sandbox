package makefriend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AddFriend struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
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
	MakeFriend(TargetID int, SourceID int) error
}

func InternalError(w http.ResponseWriter, errStr string)  {
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

	f := &AddFriend{}
	if err = json.Unmarshal(content, f); err != nil {
		InternalError(w, err.Error())
		return
	}

	err = h.groupService.MakeFriend(f.TargetID, f.SourceID)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("users %d, %d are friends", f.TargetID, f.SourceID)))
	return
}
