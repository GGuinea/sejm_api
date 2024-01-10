package api

type Video struct {
	Id            string `json:"unid"`
	Title         string `json:"title"`
	Descritpion   string `json:"descritpion"`
	Type          string `json:"type"`
	Comimitee     string `json:"comimitee"`
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
	Room          string `json:"room"`
	VideoLink     string `json:"videoLink"`
	Transcribe    bool   `json:"transcribe"`
}

func (c *Client) ListVideos() ([]Video, error) {
	url := getListVideosPath(c.URL)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	videos := []Video{}
	err = pureResponseDecoder.Decode(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (c *Client) ListTodayVideos() ([]Video, error) {
	url := getListTodayVideosPath(c.URL)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	videos := []Video{}
	err = pureResponseDecoder.Decode(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
