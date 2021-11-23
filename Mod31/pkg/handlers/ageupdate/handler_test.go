package ageupdate

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

func TestAgeUpdate(t *testing.T)  {
	user := Input{
		NewAge: 77,
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

	cl := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, srv.URL + "/619c86eaa1fb67abf8281b9a", reqBody)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	resp, err := cl.Do(req)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Log(err)
		t.Fail()
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !strings.Contains(string(respBody), "age updated") {
		t.Log(err)
		t.Fail()
	}
}
