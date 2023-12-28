package parliament

func getListVotingsPath(base, sitting string) string {
	return base + "/votings/" + sitting
}

func getVotingPath(base, sitting, votingNumber string) string {
	return base + "/votings/" + sitting + "/" + votingNumber
}
