package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"sandbox/Mod31/pkg/handlers/test"
	"testing"
)

const (
	correctResult = "User ID: 1, name: Peta was created"
)

type User struct {
	Name string
	Age int
}

func (u *User) CreateUser(name string, age int) (string, error) {
	ID := "1"
	*u = User{Name: name, Age: age}
	return fmt.Sprintf("User ID: %s, name: %s was created", ID, name), nil
}

func TestHandle_Create(t *testing.T)  {
	input := Input{
		Name: "Peta",
		Age: 15,
	}
	user := User{}

	reqBodyBytes, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)

	r := chi.NewRouter()
	r.Post("/create", NewHandler(&user))
	ts := httptest.NewServer(r)

	respStatus, respBody := test.Request(t, ts, http.MethodPost, "/create", reqBody)

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
