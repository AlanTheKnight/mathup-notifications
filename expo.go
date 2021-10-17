package main

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
)

type Notification struct {
	To    string `json:"to"`
	Sound string `json:"sound"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (n *Notification) validate() bool{
	r, _ := regexp.Compile(`ExponentPushToken\[.+]`)
	return r.MatchString(n.Title)
}

func SendPushNotification(n Notification) {
	url := "https://exp.host/--/api/v2/push/send"
	var jsonStr = []byte(``)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-encoding", "gzip, deflate")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

}
