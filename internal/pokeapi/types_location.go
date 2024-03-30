package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResp struct {
	Count     int            `json:"count"`
	Next      *string        `json:"next"`
	Previous  *string        `json:"previous"`
	Locations []LocationArea `json:"results"`
}
