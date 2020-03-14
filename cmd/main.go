package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error making a request to google :(")
		os.Exit(1)
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Couldn't read the response body :(")
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
