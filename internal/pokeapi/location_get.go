package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (DetailedLocationAreaResp, error) {
	fullUrl := baseURL + "/location-area/" + location

	fmt.Println(fullUrl)

	data, ok := c.cache.Get(fullUrl)
	if ok {
		var locationAreaResp DetailedLocationAreaResp

		if err := json.Unmarshal(data, &locationAreaResp); err != nil {
			return DetailedLocationAreaResp{}, err
		}
		fmt.Println("Cache hit!")

		return locationAreaResp, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return DetailedLocationAreaResp{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DetailedLocationAreaResp{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return DetailedLocationAreaResp{}, fmt.Errorf("bad stats code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return DetailedLocationAreaResp{}, err
	}

	var locationAreaResp DetailedLocationAreaResp

	if err = json.Unmarshal(data, &locationAreaResp); err != nil {
		return DetailedLocationAreaResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreaResp, nil
}
