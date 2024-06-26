package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(url *string) (LocationAreaResp, error) {
	fullUrl := baseURL + "/location"

	if url != nil {
		fullUrl = *url
	}

	data, ok := c.cache.Get(fullUrl)
	if ok {
		var locationAreaResp LocationAreaResp

		if err := json.Unmarshal(data, &locationAreaResp); err != nil {
			return LocationAreaResp{}, err
		}
		fmt.Println("Cache hit!")

		return locationAreaResp, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad stats code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	var locationAreaResp LocationAreaResp

	if err = json.Unmarshal(data, &locationAreaResp); err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreaResp, nil
}
