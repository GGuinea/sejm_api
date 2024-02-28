package api

import "fmt"

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"` // web-body, web-description, body
}

type Attachement struct {
	URL          string `json:"url"`
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
}

type Reply struct {
	From            string        `json:"from"`
	Key             string        `json:"key"`
	LastModified    string        `json:"lastModified"`
	ReceiptDate     string        `json:"receiptDate"`
	OnlyAttachement bool          `json:"onlyAttachement"`
	Links           []Link        `json:"links"`
	Attachements    []Attachement `json:"attachements"`
}

type Interpellation struct {
	LastModified string   `json:"lastModified"`
	Links        []Link   `json:"links"`
	Num          int      `json:"num"`
	ReceiptDate  string   `json:"receiptDate"`
	Title        string   `json:"title"`
	Term         int      `json:"term"`
	From         []string `json:"from"`
	Replies      []Reply  `json:"replies"`
	SentDate     string   `json:"sentDate"`
	To           []string `json:"to"`
}

type Params struct {
	Offset        int
	Limit         int
	SortBy        string
	From          string
	To            string
	Since         string
	Till          string
	ModifiedSince string
}

func (c *Client) ListInterpellations(params Params) ([]Interpellation, error) {
	url := getListInterpellationsPath(c.URL)
	queryParams := buildParams(params)
	if queryParams != "" {
		url += "?" + queryParams[1:]
	}

	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	interpellations := []Interpellation{}
	err = pureResponseDecoder.Decode(&interpellations)
	if err != nil {
		return nil, err
	}
	return interpellations, nil
}

func (c *Client) GetInterpelation(id string) (*Interpellation, error) {
	url := getInterpelationPath(c.URL, id)

	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	interpellation := Interpellation{}
	err = pureResponseDecoder.Decode(&interpellation)
	if err != nil {
		return nil, err
	}
	return &interpellation, nil
}

func buildParams(params Params) string {
	var result string
	if params.Offset != 0 {
		result += "&offset=" + fmt.Sprint(params.Offset)
	}
	if params.Limit != 0 {
		result += "&limit=" + fmt.Sprint(params.Limit)
	}
	if params.SortBy != "" {
		result += "&sort_by=" + params.SortBy
	}
	if params.From != "" {
		result += "&from=" + params.From
	}
	if params.To != "" {
		result += "&to=" + params.To
	}
	if params.Since != "" {
		result += "&since=" + params.Since
	}
	if params.Till != "" {
		result += "&till=" + params.Till
	}
	if params.ModifiedSince != "" {
		result += "&modifiedSince=" + params.ModifiedSince
	}
	return result
}
