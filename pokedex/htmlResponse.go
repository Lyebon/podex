package main

import (
	"net/http"
	"io"
	"encoding/json"
)

func getResponse(url *string, cfg *config) (pokemonAreas, error){
	if url == nil{
		url = &cfg.pokeApi
	}
	req, ok := cfg.cache.Get(*url)
	if !ok {
		res, err := http.Get(*url)
		if err != nil{
			return pokemonAreas{}, err
		}

		defer res.Body.Close()
		req, err = io.ReadAll(res.Body)
		if err != nil{
			return pokemonAreas{}, err
		}

		cfg.cache.Add(*url, req)
	}
	
	areas := pokemonAreas{}
	err := json.Unmarshal(req, &areas)
	if err != nil{
		return pokemonAreas{}, err
	}
	
	return areas, nil
}