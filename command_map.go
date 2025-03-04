package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type areaLocationsResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *config) mapPokedexBackward() error {
	if c.PrevURL == nil {
		return fmt.Errorf("cannot map backward, there's no previous page")
	}

	res, err := http.Get(*c.PrevURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("received non-200 status code: %d and\nbody: %s", res.StatusCode, body)
	}

	var data areaLocationsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	c.NextURL = data.Next
	c.PrevURL = data.Previous

	for _, r := range data.Results {
		fmt.Println(r.Name)
	}
	return nil
}
func (c *config) mapPokedex() error {
	res, err := http.Get(c.NextURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("received non-200 status code: %d and\nbody: %s", res.StatusCode, body)
	}

	var data areaLocationsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	c.NextURL = data.Next
	c.PrevURL = data.Previous

	for _, r := range data.Results {
		fmt.Println(r.Name)
	}
	return nil
}
