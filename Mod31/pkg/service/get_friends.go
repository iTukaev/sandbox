package service

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (d *CreateUser) Friends(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
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

		allFriendsOfUser := ""
		for _, val := range d.department.Users[uNum].Friends {
			allFriendsOfUser += strconv.Itoa(val) + "\t"
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Friends of user " + d.department.Users[uNum].Name + ": " + allFriendsOfUser))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}