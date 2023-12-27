package parliament

import ("testing")

func TestShouldAppendSittingToBaseURL(t *testing.T) {
	base := "https://api.sejm.gov.pl/sejm/term/8"
	sitting := "1"
	expected := "https://api.sejm.gov.pl/sejm/term/8/votings/1"
	actual := getVotingsPath(base, sitting)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
