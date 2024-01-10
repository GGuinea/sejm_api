package api

import (
	"testing"
)

func TestShouldReturnAllClubs(t *testing.T) {
	term := "10"
	client := NewClient(term)
	clubs, err := client.ListClubs()
	if err != nil {
		t.Error(err)
	}

	if len(clubs) == 0 {
		t.Error("No clubs returned")
	}
}

func TestShouldReturnClubById(t *testing.T) {
	term := "10"
	client := NewClient(term)
	club, err := client.GetClubById("PiS")
	if err != nil {
		t.Error(err)
	}

	if club.Id != "PiS" {
		t.Error("Wrong club returned")
	}

	if club.Name != "Klub Parlamentarny Prawo i Sprawiedliwość" {
		t.Error("Wrong club returned")
	}
}
