package deleteuser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Input struct {
	TargetID int `json:"target_id"`
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
	DeleteUser(ID int) (string, error)
}

func InternalError(w http.ResponseWriter, errStr string)  {
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

	f := &Input{}
	if err = json.Unmarshal(content, f); err != nil {
		InternalError(w, err.Error())
		return
	}

	name, err := h.groupService.DeleteUser(f.TargetID)
	if err != nil {
		InternalError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User ID: %d, Name: %s, remove successfull", f.TargetID, name )))
}
