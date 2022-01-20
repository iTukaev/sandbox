package getall

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"sandbox/Mod31/pkg/handlers/test"
	"testing"
)

const (
	correctResult = "{\"name\":\"Peta\",\"age\":15,\"friends\":[\"1\",\"2\"]}"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []string `json:"friends"`
}

func (u *User) GetAll() ([]byte, error) {
	buf, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func TestHandle_GetAll(t *testing.T)  {
	user := User{
		Name: "Peta",
		Age: 15,
		Friends: []string{"1", "2"},
	}

	r := chi.NewRouter()
	r.Get("/get", NewHandler(&user))
	ts := httptest.NewServer(r)
	defer ts.Close()

	respStatus, respBody := test.Request(t, ts, http.MethodGet, "/get", nil)

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
