package pokeapi

import (
	"fmt"
)

// GetLocationAreas fetches a list of location areas.
func (c *Client) GetLocationAreas(pageURL *string) (*ListResponse[LocationAreas], error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var response ListResponse[LocationAreas]
	if err := Fetch(c, url, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetLocation(location *string) (*Location, error) {
	url := baseURL + "/location-area"

	if location != nil {
		url = fmt.Sprintf("%s/%s", url, *location)
	}

	var response Location
	if err := Fetch(c, url, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
