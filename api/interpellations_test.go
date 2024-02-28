package api

import (
	"strings"
	"testing"
)

func TestShoudlListInterpellations(t *testing.T) {
	term := "10"
	client := NewClient(term)
	interps, err := client.ListInterpellations(Params{Limit: 1})
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(interps) == 0 {
		t.Errorf("Expected non-empty list of interpellations")
	}

	if interps[0].Title == "" {
		t.Errorf("Expected non-empty title")
	}

	if !strings.Contains(interps[0].Title, "Interpelacja") {
		t.Errorf("Expected title to contain 'Interpelacja'")
	}
}

func TestShouldReturnOneInterpelation(t *testing.T) {
	term := "10"
	client := NewClient(term)
	interpellation, err := client.GetInterpelation("1")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if interpellation == nil {
		t.Errorf("Should not be nill")
	}
	if !strings.Contains(*&interpellation.Title, "rodzinnych") {
		t.Errorf("Wrong interpellation has been return")
	}
}
