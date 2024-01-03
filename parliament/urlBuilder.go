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

func getListPrintsPath(base string) string {
	return base + "/prints"
}

func getPrintPath(base, number string) string {
	return base + "/prints/" + number
}

func getPrintAttachementPath(base, number, filename string) string {
	return base + "/prints/" + number + "/" + filename
}

func getListVideosPath(base string) string {
	return base + "/videos"
}

func getListTodayVideosPath(base string) string {
	return base + "/videos/today"
}

func getListInterpellationsPath(base string) string {
	return base + "/interpellations"
}
