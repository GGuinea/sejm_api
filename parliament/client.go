package parliament

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Term string
	URL  string
}

const (
	API_BASE_URL = "https://api.sejm.gov.pl/sejm/term"
)

func NewClient(term string) *Client {
	return &Client{Term: term, URL: API_BASE_URL + term}
}

func get(url string) ([]map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	var test []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&test)

	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %s", err)
	}

	return test, nil
}
