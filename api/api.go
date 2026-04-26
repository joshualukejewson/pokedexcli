package api

import (
	"net/http"
)

type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const LOCATION_API_URL string = "https://pokeapi.co/api/v2/location-area"

func fillConfig() (config, error) {
	res, err := http.Get(LOCATION_API_URL)
	if err != nil {
		return config{}, err
	}

}
