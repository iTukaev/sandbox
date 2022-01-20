package makefriend

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

func (u *User) MakeFriend(TargetID string, SourceID string) error {
	return nil
}

func TestHandle_GetFriends(t *testing.T)  {
	input := Input{
		TargetID: "1",
		SourceID: "2",
	}
	user := User{}

	reqBody, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}

	r := chi.NewRouter()
	r.Post("/make_friends", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodPost, "/make_friends", bytes.NewBuffer(reqBody))

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
