package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sandbox/Mod30/pkg/entity"
	"strconv"
)

type CreateUser struct {
	department *entity.Department
}

func NewCreateUser(department *entity.Department) *CreateUser {
	return &CreateUser{
		department: department,
	}
}

func (d *CreateUser) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		u := entity.NewUser()
		if err := json.Unmarshal(content, u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		uNum := lastUserId(d) + 1
		d.department.Users[uNum] = u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User ID: " + strconv.Itoa(uNum) + " name: " + u.Name + " was created"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func lastUserId(d *CreateUser) (lastId int) {
	lastId = 0
	for key,  _ := range d.department.Users {
		if key > lastId {
			lastId = key
		}
	}
	return
}