package deleteuser

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
	correctResult = "User ID: 1, was deleted"
)

type User struct {}

func (u *User) DeleteUser(ID string) (string, error) {
	return fmt.Sprintf("User ID: %s, was deleted", ID), nil
}

func TestHandle_Delete(t *testing.T)  {
	input := Input{
		TargetID: "1",
	}
	user := User{}

	reqBodyBytes, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)

	r := chi.NewRouter()
	r.Delete("/user", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodDelete, "/user", reqBody)

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
