package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

func Apiuser(user string, password string, token string) bool{

	var status bool

	data := Payload{
		Name: user,
		Token: token,
		Password: password,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {

	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://localhost:3000/api", body)
	if err != nil {

	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {

	}

	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	string_response := string(response)
	if string_response == "True" {
		status = true
	} else {
		status = false
	}
	
	return status

}