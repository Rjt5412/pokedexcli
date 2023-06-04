package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon" + "/" + pokemonName

	cachedVal, exists := c.cache.Get(url)

	if exists {
		PokemonResp := Pokemon{}
		err := json.Unmarshal(cachedVal, &PokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return PokemonResp, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	PokemonResp := Pokemon{}
	err = json.Unmarshal(data, &PokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return PokemonResp, nil
}
