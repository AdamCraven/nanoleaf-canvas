package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getAuthToken() (string, string) {
	nanoleafIP := "192.168.8.170:16021"
	url := "http://" + nanoleafIP + "/api/v1/new"

	req, err := http.NewRequest("POST", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status, string(body)
}

func main() {
	status, body := getAuthToken()

	if status != "200 OK" {
		fmt.Println("Hold power button on canvas for around 7 seconds to get auth token...")

		time.Sleep(2000 * time.Millisecond)
		main()
	} else {
		fmt.Println("response Body:", body)
	}
}
