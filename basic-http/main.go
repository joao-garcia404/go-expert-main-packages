package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	// This line will be executed after all instructions (defer)
	defer request.Body.Close()

	// Reading the http stream
	response, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	println(string(response))
}
