package deleteuser

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
		TargetID: "619c850c5d3839cb4abab143",
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
	req, err := http.NewRequest(http.MethodDelete, srv.URL, reqBody)
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

	if !strings.Contains(string(respBody), "was deleted") {
		t.Log(err)
		t.Fail()
	}
}
