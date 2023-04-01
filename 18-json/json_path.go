package main

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
)

type JsonPathExample struct {
	custom interface{}
}

func (j *JsonPathExample) Run() {
	fmt.Println("Hi! JsonPath started")
	j.custom = interface{}(nil)

	json.Unmarshal([]byte(`{
		"welcome":{
				"message":["Good Morning", "Hello World!"]
			}
		}`), &j.custom)

	welcome, err := jsonpath.Get("$.welcome.message[1]", j.custom)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(welcome)
}
