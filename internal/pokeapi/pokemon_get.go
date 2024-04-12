package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	fullUrl := baseURL + "/pokemon/" + name

	data, ok := c.cache.Get(fullUrl)
	if ok {
		var pokemon Pokemon

		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		fmt.Println("Cache hit!")

		return pokemon, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad stats code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon

	if err = json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemon, nil
}
