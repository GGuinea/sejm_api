package parliament

import (
	"strings"
	"testing"
)

func TestShouldListVideos(t *testing.T) {
	term := "10"
	client := NewClient(term)
	videos, err := client.ListVideos()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(videos) == 0 {
		t.Errorf("Expected non-empty list of videos")
	}

	if videos[0].Title == "" {
		t.Errorf("Expected non-empty title")
	}

	if !strings.Contains(videos[0].Title, "Planowane posiedzenie") {
		t.Errorf("Expected title to contain 'Planowane posiedzenie'")
	}
}

func TestShouldListTodayVideos(t *testing.T) {
	term := "10"
	client := NewClient(term)
	_, err := client.ListTodayVideos()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
