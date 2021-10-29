package service

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"sandbox/Mod31/pkg/entity"
	"strconv"
)

func (d *CreateUser) AgeUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		router := chi.URLParam(r, "userId")

		uNum, err := strconv.Atoi(router)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("incorrect user ID"))
			return
		}

		if _, ok := d.department.Users[uNum]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user ID not found"))
			return
		}

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		u := entity.NewUpdateAge()
		if err := json.Unmarshal(content, u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		d.department.Users[uNum].Age = u.NewAge

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Age of user " + d.department.Users[uNum].Name + " updated successful"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}