package parliament

import (
	"testing"
)

func TestShouldReturnAllCommittees(t *testing.T) {
	term := "10"
	client := NewClient(term)
	committees, err := client.ListCommittees()
	if err != nil {
		t.Error(err)
	}

	if len(committees) == 0 {
		t.Error("No committees returned")
	}
}
