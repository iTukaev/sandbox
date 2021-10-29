package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sandbox/Mod31/pkg/entity"
)

func (d *CreateUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
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

		if _, ok := d.department.Users[f.TargetId]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user ID not found"))
			return
		}
		name := d.department.Users[f.TargetId].Name
		delete(d.department.Users, f.TargetId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User " + name +	" deleted"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}