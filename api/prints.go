package api

type Print struct {
	Number           string   `json:"number"`
	Term             int      `json:"term"`
	Title            string   `json:"title"`
	ChangeDate       string   `json:"changeDate"`
	DeliveryDate     string   `json:"deliveryDate"`
	DocumentDate     string   `json:"documentDate"`
	ProcessPrint     []string `json:"processPrint"`
	Attachements     []string `json:"attachements"`
	AdditionalPrints []Print  `json:"additionalPrints"`
	NumberAssociated []string `json:"numberAssociated"`
}

func (c *Client) ListPrints() ([]Print, error) {
	url := getListPrintsPath(c.URL)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	prints := []Print{}
	err = pureResponseDecoder.Decode(&prints)
	if err != nil {
		return nil, err
	}
	return prints, nil
}

func (c *Client) GetPrint(number string) (*Print, error) {
	url := getPrintPath(c.URL, number)
	pureResponseDecoder, err := get(url)
	if err != nil {
		return nil, err
	}

	print := Print{}
	err = pureResponseDecoder.Decode(&print)
	if err != nil {
		return nil, err
	}
	return &print, nil
}

func (c *Client) GetPrintAttachement(number, filename string) (*[]byte, error) {
	url := getPrintAttachementPath(c.URL, number, filename)
	attachement, err := getFile(url)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &attachement, nil
}
