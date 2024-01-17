package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Term string
	URL  string
}

const (
	API_BASE_URL = "http://api.sejm.gov.pl/sejm/term"
)

func NewClient(term string) *Client {
	return &Client{Term: term, URL: API_BASE_URL + term}
}

func get(url string) (*json.Decoder, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	return decoder, nil
}

func getFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	return body, nil
}
