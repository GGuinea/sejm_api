package parliament

import (
	"testing"
)

func TestNewClientShouldReturnNewClinetStructWithParamsSet(t *testing.T) {
	term := "8"
	client := NewClient(term)
	if client.Term != term {
		t.Error("Expected term to be set to 8")
	}
	if client.URL != API_BASE_URL+term {
		t.Error("Expected URL to be set to API_BASE_URL + 8")
	}
}
