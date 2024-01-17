package api

type Club struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	MemgerCount int    `json:"memberCount"`
}

func (c *Client) ListClubs() (clubs []Club, err error) {
	url := getListClubsPath(c.URL)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	clubs = make([]Club, 0)
	err = pureResponseDecoder.Decode(&clubs)
	if err != nil {
		return nil, err
	}

	return clubs, nil
}

func (c *Client) GetClubById(id string) (club Club, err error) {
	url := getClubPath(c.URL, id)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return club, err
	}

	err = pureResponseDecoder.Decode(&club)
	if err != nil {
		return club, err
	}

	return club, nil
}
