package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location?offset=0&limit=20"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationAreaResp := LocationAreaResp{}

		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
	}
	fmt.Println("cache miss !")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}

	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationAreaResp, nil
}
