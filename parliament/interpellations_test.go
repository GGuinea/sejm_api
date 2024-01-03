package parliament

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
