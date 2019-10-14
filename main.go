package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var apiKey string
var networkAddress string

var plugin string = "b3fd723a-aae8-4c99-bf2b-087159e0ef53"

func main() {
	apiKey = os.Getenv("nanoleaf_canvas_api")
	networkAddress = os.Getenv("nanoleaf_canvas_ip")

	if apiKey == "" {
		log.Fatal("nanoleaf_canvas_api env variable not setup")
	}
	if networkAddress == "" {
		log.Fatal("nanoleaf_canvas_ip env variable not setup")
	}

	fmt.Printf("Network address is: %s. apiKey starts with: %s...", networkAddress, apiKey[:5])

	lightOn()
	time.Sleep(3 * time.Second)
	animate()
	time.Sleep(3 * time.Second)

	lightOff()
}

func requestPut(payload *strings.Reader, endpoint string) {

	apiURL := fmt.Sprintf("http://%s:16021/api/v1/", networkAddress)
	requestURL := apiURL + apiKey + endpoint

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
	requestPut(payload, "/state")
}

func lightOff() {
	payload := strings.NewReader("{\"on\" : {\"value\":false}}")
	requestPut(payload, "/state")
}

func animate() {
	payload := strings.NewReader("{\"write\":{\"command\":\"display\",\"animType\":\"static\",\"animData\":\"3 82 1 255 0 255 0 20 60 1 0 255 255 0 20 118 1 0 0 0 0 20\",\"loop\":false}}")
	requestPut(payload, "/effects")

}

/*
	"{\"write\":{\"command\":\"add\",\"version\":\"2.0\",\"animType\":\"plugin\",\"animName\":\"MyAnimation\",\"colorType\":\"HSB\",\"pluginUuid\":\"b3fd723a-aae8-4c99-bf2b-087159e0ef53\",\"pluginType\":\"color\",\"pluginOptions\":[{\"name\":\"transTime\",\"value\":2},{\"name\":\"direction\",\"value\":\"left\"},{\"name\":\"loop\",\"value\":true}],\"Palette\":[{\"hue\":0,\"saturation\":100,\"brightness\":100},{\"hue\":120,\"saturation\":100,\"brightness\":100},{\"hue\":240,\"saturation\":100,\"brightness\":100}]}}")*/
