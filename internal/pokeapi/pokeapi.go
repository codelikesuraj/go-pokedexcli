package pokeapi

const baseURL = "https://pokeapi.co/api/v2"

func GetBaseURL() *string {
	url := baseURL
	return &url
}
