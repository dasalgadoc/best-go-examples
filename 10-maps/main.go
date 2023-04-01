package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["diego"] = 100
	myMap["daniela"] = 200

	fmt.Println(myMap)

	for i, value := range myMap {
		fmt.Println(i, value)
	}

	value := myMap["daniela"]
	fmt.Println(value)

	value = myMap["ignacio"]
	fmt.Println(value)

	value, ok := myMap["ignacio"]
	fmt.Println(ok, value)

	if !ok {
		fmt.Printf("Value %s does not exists\n", "ignacio")
	}

	value, ok = myMap["daniela"]
	fmt.Println(ok, value)

	if !ok {
		fmt.Printf("Value %s does not exists", "daniela")
	}

}
