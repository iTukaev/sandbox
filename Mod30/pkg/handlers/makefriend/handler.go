package makefriend

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AddFriend struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type Handle struct {
	GroupService groupInterface
}

type groupInterface interface {
	MakeFriend(TargetID int, SourceID int) (string, error)
}

func (h *Handle) MakeFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		f := &AddFriend{}
		if err := json.Unmarshal(content, f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		responseBody, err := h.GroupService.MakeFriend(f.TargetID, f.SourceID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseBody))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
