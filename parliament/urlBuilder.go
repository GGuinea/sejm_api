package parliament

func getListVotingsPath(base, sitting string) string {
	return base + "/votings/" + sitting
}

func getVotingPath(base, sitting, votingNumber string) string {
	return base + "/votings/" + sitting + "/" + votingNumber
}

func getListClubsPath(base string) string {
	return base + "/clubs"
}

func getClubPath(base, id string) string {
	return base + "/clubs/" + id
}

func getListCommitteesPath(base string) string {
	return base + "/committees"
}

func getListEnvoysPath(base string) string {
	return base + "/MP"
}

func getEnvoyPath(base, id string) string {
	return base + "/MP/" + id
}
