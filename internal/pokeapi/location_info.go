package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationInfo(area *string) (RespLocationsInfo, error) {
	url := baseURL + "/location-area/" + *area

	if dat, ok := c.cache.Get(url); ok {
		locationInfoResp := RespLocationsInfo{}
		err := json.Unmarshal(dat, &locationInfoResp)
		if err != nil {
			return RespLocationsInfo{}, err
		}
		return locationInfoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsInfo{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationsInfo{}, err
	}

	locationsInfoResp := RespLocationsInfo{}
	err = json.Unmarshal(dat, &locationsInfoResp)
	if err != nil {
		return RespLocationsInfo{}, err
	}

	c.cache.Add(url, dat)
	return locationsInfoResp, nil
}
