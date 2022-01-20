package deleteuser

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
	correctResult = "User ID 0000 was deleted"
)

type User struct {}

func (u *User) DeleteUser(ID string) error {
	return nil
}

func TestHandle_Delete(t *testing.T)  {
	input := Input{
		TargetID: "0000",
	}
	user := User{}

	reqBody, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}

	r := chi.NewRouter()
	r.Delete("/user", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodDelete, "/user", bytes.NewBuffer(reqBody))

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
