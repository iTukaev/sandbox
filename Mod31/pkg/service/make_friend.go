package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sandbox/Mod31/pkg/entity"
	"strconv"
)

func (d *CreateUser) MakeFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		f := entity.NewAddFriend()
		if err := json.Unmarshal(content, f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := findFriend(&d.department.Users[f.TargetId].Friends, f.SourceId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		d.department.Users[f.TargetId].Friends = append(d.department.Users[f.TargetId].Friends, f.SourceId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User ID: " + strconv.Itoa(f.SourceId) +
			" added to friends for user ID: " + strconv.Itoa(f.TargetId)))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func findFriend(s *[]int, num int) error {
	var err error = nil

	for _, val := range *s {
		if val == num {
			return errors.New("these users are already friends")
		}
	}
	return err
}