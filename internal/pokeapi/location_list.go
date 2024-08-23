package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c  *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil { return RespLocations{}, err }
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil { return RespLocations{}, err }

	resp, err := c.httpClient.Do(req)
	if err != nil { return RespLocations{}, err }
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil { return RespLocations{}, err }

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil { return RespLocations{}, err}

	c.cache.Add(url, data)
	return locationsResp, nil
}