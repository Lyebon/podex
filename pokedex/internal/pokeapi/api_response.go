package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaResponse(pageURL *string) (pokeAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return pokeAreas{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return pokeAreas{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return pokeAreas{}, fmt.Errorf("status code %v", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return pokeAreas{}, err
		}

		c.cache.Add(url, data)
	}

	areas := pokeAreas{}
	err := json.Unmarshal(data, &areas)
	if err != nil {
		return pokeAreas{}, err
	}

	return areas, nil
}


// refactorizar despues cuando funcione explore
func (c *Client) GetPokemonEncounter(pageURL string) (pokeEncounters, error) {
	url := baseURL + "/location-area/" + pageURL
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return pokeEncounters{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return pokeEncounters{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return pokeEncounters{}, fmt.Errorf("status code %v", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return pokeEncounters{}, err
		}

		c.cache.Add(url, data)
	}

	pokemons := pokeEncounters{}
	err := json.Unmarshal(data, &pokemons)
	if err != nil {
		return pokeEncounters{}, err
	}

	return pokemons, nil

}