package getfriends

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"sandbox/Mod31/pkg/handlers/test"
	"testing"
)

const (
	correctResult = "Friends of user 0000 - [1 2]"
)

type User struct {
	Friends []string
}

func (u *User) GetFriends(ID string) ([]string, error) {
	return u.Friends, nil
}

func TestHandle_GetFriends(t *testing.T)  {
	userID := "0000"
	user := User{
		Friends: []string{"1", "2"},
	}

	r := chi.NewRouter()
	r.Get("/friends/{userId}", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodGet, "/friends/" + userID, nil)

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
