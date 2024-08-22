package pokeapi

type RespLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Result []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}