package create

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"sandbox/Mod31/pkg/handlers/test"
	"testing"
)

const (
	correctResult = "User name: Peta created with ID: 0000"
)

type User struct {
	Name string
	Age int
}

func (u *User) CreateUser(name string, age int) (string, error) {
	ID := "0000"
	u.Name = name
	u.Age = age
	return ID, nil
}

func TestHandle_Create(t *testing.T)  {
	input := Input{
		Name: "Peta",
		Age: 15,
	}
	user := User{}

	reqBody, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}

	r := chi.NewRouter()
	r.Post("/create", NewHandler(&user))
	ts := httptest.NewServer(r)

	respStatus, respBody := test.Request(t, ts, http.MethodPost, "/create", bytes.NewBuffer(reqBody))

	if respStatus != http.StatusCreated {
		t.Log("status fail")
		t.Fail()
		return
	}

	if respBody != correctResult {
		t.Log("body fail")
		t.Fail()
	}
}
