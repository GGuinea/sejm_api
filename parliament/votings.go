package parliament

type Voting struct {
	Abstain          int    `json:"abstain"`
	Date             string `json:"date"`
	Description      string `json:"description"`
	Kind             string `json:"kind"`
	No               int    `json:"no"`
	NotParticipating int    `json:"notParticipating"`
	Sitting          int    `json:"sitting"`
	SittingDay       int    `json:"sittingDay"`
	Term             int    `json:"term"`
	Title            string `json:"title"`
	Topic            string `json:"topic"`
	TotalVoted       int    `json:"totalVoted"`
	VotingNumber     int    `json:"votingNumber"`
	VotingOptions    []VotingOption
	Votes            []Vote `json:"votes"`
}

type VotingOption struct {
	Option      string `json:"option"`
	OptionIndex int    `json:"optionIndex"`
	Votes       int    `json:"votes"`
}

type Vote struct {
	MP         int               `json:"MP"`
	Club       string            `json:"club"`
	FirstName  string            `json:"firstName"`
	LastName   string            `json:"lastName"`
	SecondName string            `json:"secondName"`
	ListVotes  map[string]string `json:"listVotes"`
	Vote       string            `json:"vote"`
}

func (client *Client) ListVotings(sitting string) ([]Voting, error) {
	url := getListVotingsPath(client.URL, sitting)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	votingResponse := make([]Voting, 0)
	err = pureResponseDecoder.Decode(&votingResponse)
	if err != nil {
		return nil, err
	}

	return votingResponse, nil
}

func (client *Client) GetVoting(sitting, votingNumber string) (*Voting, error) {
	url := getVotingPath(client.URL, sitting, votingNumber)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	votingResponse := Voting{}
	err = pureResponseDecoder.Decode(&votingResponse)
	if err != nil {
		return nil, err
	}

	return &votingResponse, nil
}
