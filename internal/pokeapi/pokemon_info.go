package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(name *string) (RespPokemonInfo, error) {
	url := baseURL + "/pokemon/" + *name

	if dat, ok := c.cache.Get(url); ok {
		pokemonInfo := RespPokemonInfo{}
		err := json.Unmarshal(dat, &pokemonInfo)
		if err != nil {
			return RespPokemonInfo{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	if res.StatusCode >= 400 {
		return RespPokemonInfo{}, fmt.Errorf("%v", res.StatusCode)
	}
	pokemonInfo := RespPokemonInfo{}
	err = json.Unmarshal(dat, &pokemonInfo)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	c.cache.Add(url, dat)
	return pokemonInfo, nil
}
