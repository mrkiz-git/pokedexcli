package pokeapi

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// ListLocationAreas returns a list of location areas from the PokeAPI
func (c *Client) ListLocationAreas(pageURL *string) (*ListResponse[LocationAreas], error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ListResponse[LocationAreas]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
