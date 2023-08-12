package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonList struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonStats struct {
	Height int `json:"height"`
}

func GetPokemonList() (data PokemonList) {
	url := "https://pokeapi.co/api/v2/pokemon?limit=300"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error on:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error JSON Unmarshall:", err)
		return
	}

	return
}

func GetPokemonHeight(pokemon string) (stats PokemonStats) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error on:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	err = json.Unmarshal(body, &stats)
	if err != nil {
		fmt.Println("Error JSON Unmarshall:", err)
		return
	}

	//log.Println("Pokemon", pokemon, "height is", stats.Height)

	return
}
