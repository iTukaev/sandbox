package create

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sandbox/Mod31/pkg/db/dbService"
	"strings"
	"testing"
)

func TestHandle_Create(t *testing.T)  {
	user := Input{
		Name: "Peta",
		Age: 15,
	}
	reqBodyBytes, err := json.Marshal(&user)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)

	client := dbService.NewService()
	srv := httptest.NewServer(http.HandlerFunc(NewHandler(client)))

	resp, err := http.Post(srv.URL, "application/json", reqBody)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		t.Log(err)
		t.Fail()
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !strings.Contains(string(respBody), "Peta") {
		t.Log(err)
		t.Fail()
	}
}
