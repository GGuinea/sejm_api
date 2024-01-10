package api

func getListVotingsPath(base, sitting string) string {
	return base + "/votings/" + sitting
}

func getVotingPath(base, sitting, votingNumber string) string {
	return getListVotingsPath(base, sitting) + "/" + votingNumber
}

func getListClubsPath(base string) string {
	return base + "/clubs"
}

func getClubPath(base, id string) string {
	return getListClubsPath(base) + "/" + id
}

func getListCommitteesPath(base string) string {
	return base + "/committees"
}

func getListEnvoysPath(base string) string {
	return base + "/MP"
}

func getEnvoyPath(base, id string) string {
	return getListEnvoysPath(base) + "/" + id
}

func getListPrintsPath(base string) string {
	return base + "/prints"
}

func getPrintPath(base, number string) string {
	return getListPrintsPath(base) + "/" + number
}

func getPrintAttachementPath(base, number, filename string) string {
	return getPrintPath(base, number) + "/" + filename
}

func getListVideosPath(base string) string {
	return base + "/videos"
}

func getListTodayVideosPath(base string) string {
	return getListVideosPath(base) + "/today"
}

func getListInterpellationsPath(base string) string {
	return base + "/interpellations"
}
