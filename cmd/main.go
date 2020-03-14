package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getResponseBody(req *http.Request) (*string, error) {
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)
	return &bodyString, nil
}

func main() {
	fmt.Println("Hello world!")
	req, err := http.NewRequest(http.MethodGet, "http://google.com", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := getResponseBody(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*body)
}
