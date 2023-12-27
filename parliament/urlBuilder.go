package parliament

func getVotingsPath(base, sitting string) string {
	return base + "/votings/" + sitting
}
