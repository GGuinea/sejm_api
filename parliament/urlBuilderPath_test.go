package parliament

import (
	"testing"
)

func TestShouldAppendSittingToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	sitting := "1"
	expected := "https://api.sejm.gov.pl/sejm/term/8/votings/1"
	actual := getListVotingsPath(base, sitting)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendSittingAndVotingNumberToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	sitting := "1"
	votingNumber := "1"
	expected := "https://api.sejm.gov.pl/sejm/term/8/votings/1/1"
	actual := getVotingPath(base, sitting, votingNumber)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendClubsToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	expected := "https://api.sejm.gov.pl/sejm/term/8/clubs"
	actual := getListClubsPath(base)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendClubIdToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	id := "PiS"
	expected := "https://api.sejm.gov.pl/sejm/term/8/clubs/PiS"
	actual := getClubPath(base, id)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendCommitteesToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	expected := "https://api.sejm.gov.pl/sejm/term/8/committees"
	actual := getListCommitteesPath(base)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendEnvoysToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	expected := "https://api.sejm.gov.pl/sejm/term/8/MP"
	actual := getListEnvoysPath(base)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendEnvoyIdToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	id := "123"
	expected := "https://api.sejm.gov.pl/sejm/term/8/MP/123"
	actual := getEnvoyPath(base, id)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendPrintsToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	expected := "https://api.sejm.gov.pl/sejm/term/8/prints"
	actual := getListPrintsPath(base)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendPrintNumberToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	number := "123"
	expected := "https://api.sejm.gov.pl/sejm/term/8/prints/123"
	actual := getPrintPath(base, number)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendPrintNumberAndFilenameToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	number := "123"
	filename := "123.pdf"
	expected := "https://api.sejm.gov.pl/sejm/term/8/prints/123/123.pdf"
	actual := getPrintAttachementPath(base, number, filename)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendVideosToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	expected := "https://api.sejm.gov.pl/sejm/term/8/videos"
	actual := getListVideosPath(base)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestShouldAppendTodayVideosToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	expected := "https://api.sejm.gov.pl/sejm/term/8/videos/today"
	actual := getListTodayVideosPath(base)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
