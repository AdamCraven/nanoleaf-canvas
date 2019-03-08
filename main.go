package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	lightOn()
	time.Sleep(3 * time.Second)
	lightOff()
}

func requestPut(payload *strings.Reader) {
	apiURL := "http://192.168.8.170:16021/api/v1/"
	authToken := "xxx"
	requestURL := apiURL + authToken + "/state"

	client := &http.Client{}
	request, err := http.NewRequest("PUT", requestURL, payload)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(requestURL, payload, response.StatusCode)
	header := response.Header
	for key, value := range header {
		fmt.Println(key, ":", value)
	}
	fmt.Println(content)
}

func lightOn() {
	payload := strings.NewReader("{\"on\" : {\"value\":true}}")
	requestPut(payload)
}

func lightOff() {
	payload := strings.NewReader("{\"on\" : {\"value\":false}}")
	requestPut(payload)
}
