package ageupdate

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
	correctResult = ""
)

type User struct {}

func (u *User) AgeUpdate(ID string, age int) error {
	return nil
}

func TestHandle_AgeUpdate(t *testing.T)  {
	input := Input{
		NewAge: 77,
	}

	user := User{}

	reqBody, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}

	r := chi.NewRouter()
	r.Put("/{userId}", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodPut, "/1", bytes.NewBuffer(reqBody))

	if respStatus != http.StatusNoContent {
		t.Log("status fail")
		t.Fail()
		return
	}

	if respBody != correctResult {
		t.Log("body fail")
		t.Fail()
	}
}