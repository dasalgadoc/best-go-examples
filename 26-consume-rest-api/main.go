package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const URL = "https://pokeapi.co/api/v2/pokemon/ditto"

func main() {
	fmt.Println("basic usage")
	basicUsage()
	fmt.Println("=====")

	fmt.Println("with timeout")
	withTimeOut()
	fmt.Println("=====")

	fmt.Println("setting method")
	settingMethod()
	fmt.Println("=====")

	fmt.Println("read response")
	readResponse()
	fmt.Println("=====")

	fmt.Println("response to var")
	responseToVar()
	fmt.Println("=====")
}

func responseToVar() {
	type response map[string]interface{}
	var data response
	resp, _ := http.Get(URL)
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func readResponse() {
	resp, _ := http.Get(URL)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func settingMethod() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, _ := http.NewRequest("GET", URL, nil)
	resp, _ := client.Do(req)
	fmt.Println(resp)
}

func withTimeOut() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, _ := client.Get(URL)
	fmt.Println(resp)
}

func basicUsage() {
	resp, _ := http.Get(URL)
	fmt.Println(resp)
}
