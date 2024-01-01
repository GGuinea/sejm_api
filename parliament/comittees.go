package parliament

type Member struct {
	Club          string `json:"club"`
	Id            int    `json:"id"`
	LastFirstName string `json:"lastFirstName"`
}

type Committee struct {
	Name            string   `json:"name"`
	Scope           string   `json:"scope"`
	AppointmentDate string   `json:"appointmentDate"`
	Code            string   `json:"code"`
	CompositionDate string   `json:"compositionDate"`
	Function        string   `json:"function"`
	Members         []Member `json:"members"`
}

func (c *Client) ListCommittees() (committees []Committee, err error) {
	url := getListCommitteesPath(c.URL)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	committees = make([]Committee, 0)
	err = pureResponseDecoder.Decode(&committees)
	if err != nil {
		return nil, err
	}

	return committees, nil
}
