package api

import (
	"strings"
	"testing"
)

func TestShouldGetPrintsList(t *testing.T) {
	term := "10"
	client := NewClient(term)
	prints, err := client.ListPrints()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(prints) == 0 {
		t.Errorf("Expected non-empty list of prints")
	}

	if prints[0].Title == "" {
		t.Errorf("Expected non-empty title")
	}

	if !strings.Contains(prints[0].Title, "Poselski projekt uchwały") {
		t.Errorf("Expected title to contain 'Poselski projekt uchwały'")
	}
}

func TestShouldGetPrint(t *testing.T) {
	term := "10"
	client := NewClient(term)
	print, err := client.GetPrint("1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if print.Title == "" {
		t.Errorf("Expected non-empty title")
	}

	if !strings.Contains(print.Title, "Poselski projekt uchwały") {
		t.Errorf("Expected title to contain 'Poselski projekt uchwały'")
	}
}

func TestShouldGetPrintAttachement(t *testing.T) {
	term := "9"
	client := NewClient(term)
	attachement, err := client.GetPrintAttachement("383", "383.pdf")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(*attachement) == 0 {
		t.Errorf("Expected non-empty attachement")
	}
}
