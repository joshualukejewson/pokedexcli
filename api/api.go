package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const LOCATION_API_URL string = "https://pokeapi.co/api/v2/location-area"

func FillConfig() (Config, error) {
	res, err := http.Get(LOCATION_API_URL)
	if err != nil {
		return Config{}, err
	}

	locationData, err := io.ReadAll(res.Body)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	err = json.Unmarshal(locationData, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
