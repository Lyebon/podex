package pokeapi

type pokeAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokeEncounters struct {
	Name			 string	`json:"name"`
	PokemonEncounter []map[string]string `json:"pokemon_encounters"`
}
