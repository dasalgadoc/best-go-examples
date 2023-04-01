package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type config struct {
	IO      io       `yaml:"io"`
	Clients []client `yaml:"clients"`
}

type io struct {
	FolderLocation string `yaml:"location"`
	InputFileName  string `yaml:"file_name"`
}

type client struct {
	Name    string            `yaml:"name"`
	Type    string            `yaml:"type"`
	Path    string            `yaml:"path"`
	Headers map[string]string `yaml:"headers"`
}

func main() {
	entryString := getString()
	data := []byte(entryString)

	var myConfig config
	err := yaml.Unmarshal(data, &myConfig)

	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("%+v\n", myConfig)
}

func getString() string {
	return `io:
  location: ./file_exchange/
  file_name: input_capswrapper
clients:
  - name: api-engine
    type: GetRestApi
    path: http://internal-api.mercadopago.com/transactions/regulations/status
    headers:
      X-Client-Id: 6110093112793568
      X-Version: v2
  - name: api-engine_w
    type: GetRestApi
    path: http://internal-api.mercadopago.com/cs/withdrawals/regulations/status
    headers:
      X-Client-Id: 7707130789914454`
}
