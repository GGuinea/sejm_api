package api

type Envoy struct {
	Active         bool   `json:"active"`
	BirthDate      string `json:"birthDate"`
	BirthLocation  string `json:"birthLocation"`
	Club           string `json:"club"`
	DistrictName   string `json:"districtName"`
	DistrictNumber int    `json:"districtNum"`
	EducationLevel string `json:"educationLevel"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Id             int    `json:"id"`
	NumberOfVotes  int    `json:"numberOfVotes"`
	Voivodeship    string `json:"voivodeship"`
}

func (c *Client) ListEnvoys() (envoys []Envoy, err error) {
	url := getListEnvoysPath(c.URL)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	envoys = make([]Envoy, 0)
	err = pureResponseDecoder.Decode(&envoys)
	if err != nil {
		return nil, err
	}

	return envoys, nil
}

func (c *Client) GetEnvoyById(id string) (envoy Envoy, err error) {
	url := getEnvoyPath(c.URL, id)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return envoy, err
	}

	err = pureResponseDecoder.Decode(&envoy)
	if err != nil {
		return envoy, err
	}

	return envoy, nil
}
