package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetNext(cfg *Config) error {
	newConfig := Config{}
	res, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	locationData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(locationData, &newConfig)
	if err != nil {
		return err
	}
	*cfg = newConfig
	return nil
}

func GetPrevious(cfg *Config) error {

	if cfg.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}

	newConfig := Config{}
	res, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}
	locationData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(locationData, &newConfig)
	if err != nil {
		return err
	}
	*cfg = newConfig
	return nil
}
