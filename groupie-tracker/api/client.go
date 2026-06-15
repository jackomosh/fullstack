package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

type Client struct {
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// FetchData requests artists and relations, merging them into a UnifiedRegistry cache map.
func (c *Client) FetchData() (*UnifiedRegistry, error) {
	artists, err := c.FetchArtists()
	if err != nil {
		return nil, fmt.Errorf("artists fetch error: %w", err)
	}

	relationsMap, err := c.FetchRelations()
	if err != nil {
		return nil, fmt.Errorf("relations fetch error: %w", err)
	}

	return &UnifiedRegistry{
		Artists:   artists,
		Relations: relationsMap,
	}, nil
}

func (c *Client) FetchArtists() ([]Artist, error) {
	resp, err := c.HTTPClient.Get(BaseURL + "/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func (c *Client) FetchRelations() (map[int]Relation, error) {
	resp, err := c.HTTPClient.Get(BaseURL + "/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var wrapper RelationIndex
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}

	relationsMap := make(map[int]Relation)
	for _, rel := range wrapper.Index {
		relationsMap[rel.ID] = rel
	}
	return relationsMap, nil
}