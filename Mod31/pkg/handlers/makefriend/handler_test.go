package makefriend

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
	correctResult = "User ID: 2 now friend to user ID: 1"
)

type User struct {}

func (u *User) MakeFriend(TargetID string, SourceID string) (string, error) {
	return fmt.Sprintf("User ID: %s now friend to user ID: %s", SourceID, TargetID), nil
}

func TestHandle_GetFriends(t *testing.T)  {
	input := Input{
		TargetID: "1",
		SourceID: "2",
	}
	user := User{}

	reqBodyBytes, err := json.Marshal(&input)
	if err != nil {
		t.Fatal(err)
		return
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)

	r := chi.NewRouter()
	r.Post("/make_friends", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodPost, "/make_friends", reqBody)

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
