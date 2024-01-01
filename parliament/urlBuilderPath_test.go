package parliament

import ("testing")

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
