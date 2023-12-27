package parliament

func (client *Client) ListVotings(sitting string) {
	url := getVotingsPath(client.URL, sitting)
	get(url)
}
