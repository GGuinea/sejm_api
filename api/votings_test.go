package api

import (
	"testing"
)

func TestShouldReturnAllVotings(t *testing.T) {
	term := "10"
	sitting := "1"
	client := NewClient(term)
	votings, err := client.ListVotings(sitting)
	if err != nil {
		t.Error(err)
	}
	if len(votings) == 0 {
		t.Error("Expected votings to be returned")
	}
}

func TestShouldReturnVoting(t *testing.T) {
	term := "10"
	sitting := "1"
	votingNumber := "1"
	client := NewClient(term)
	voting, err := client.GetVoting(sitting, votingNumber)
	if err != nil {
		t.Error(err)
	}
	if voting == nil {
		t.Error("Expected voting to be returned")
	}

	if voting.Title != "Wybór Marszałka Sejmu RP" {
		t.Error("Expected voting title to be Wybór Marszałka Sejmu RP")
	}

	if len(voting.Votes) == 0 {
		t.Error("Expected votes to be returned")
	}
}
