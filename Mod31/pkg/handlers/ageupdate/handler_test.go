package ageupdate

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
	correctResult = "User's ID: 1 age updated"
)

type User struct {}

func (u *User) AgeUpdate(ID string, age int) (string, error) {
	return fmt.Sprintf("User's ID: %s age updated", ID), nil
}

func TestHandle_AgeUpdate(t *testing.T)  {
	input := Input{
		NewAge: 77,
	}

	user := User{}

	reqBodyBytes, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)

	r := chi.NewRouter()
	r.Put("/{userId}", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodPut, "/1", reqBody)

	if respStatus != http.StatusOK {
		t.Log("status fail")
		t.Fail()
		return
	}

	if respBody != correctResult {
		t.Log("body fail")
		t.Fail()
	}
}