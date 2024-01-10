package api

import (
	"testing"
)

func TestShouldReturnAllEnvoys(t *testing.T) {
	term := "10"
	client := NewClient(term)
	envoys, err := client.ListEnvoys()
	if err != nil {
		t.Error(err)
	}

	if len(envoys) == 0 {
		t.Error("No envoys returned")
	}
}

func TestShouldReturnEnvoyById(t *testing.T) {
	term := "10"
	client := NewClient(term)
	envoy, err := client.GetEnvoyById("1")
	if err != nil {
		t.Error(err)
	}

	if envoy.Id != 1 {
		t.Error("Wrong envoy returned")
	}

	if envoy.FirstName != "Andrzej" {
		t.Error("Wrong envoy returned")
	}
}
