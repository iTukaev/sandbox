package getfriends

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"sandbox/Mod31/pkg/handlers/test"
	"testing"
)

const (
	correctResult = "Friends of user ID: 1 - [1 2]"
)

type User struct {
	Friends []string
}

func (u *User) GetFriends(ID string) (string, error) {
	return fmt.Sprintf("Friends of user ID: %s - %v", ID, u.Friends), nil
}

func TestHandle_GetFriends(t *testing.T)  {
	user := User{
		Friends: []string{"1", "2"},
	}

	r := chi.NewRouter()
	r.Get("/friends/{userId}", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodGet, "/friends/1", nil)

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
