package deleteuser

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
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
	DeleteUser(ID string) (string, error)
}

func (h *Handle) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	f := &Input{}
	if err := json.Unmarshal(content, f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := h.groupService.DeleteUser(f.TargetID)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
