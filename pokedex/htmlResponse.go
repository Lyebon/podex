package main

import (
	"net/http"
	"encoding/json"
)

func getResponse(url *string) (pokemonAreas, error){
	urlApi:= "https://pokeapi.co/api/v2/location-area/"
	if url != nil{
		urlApi = *url
	}
	res, err := http.Get(urlApi)
	if err != nil{
		return pokemonAreas{}, err
	}

	defer res.Body.Close()

	areas := pokemonAreas{}
	err = json.NewDecoder(res.Body).Decode(&areas)
	if err != nil{
		return pokemonAreas{}, err
	}
	
	return areas, nil
}