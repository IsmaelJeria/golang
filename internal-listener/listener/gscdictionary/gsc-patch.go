package gscdictionary

//GSCPatchOrigin ...
type GSCPatchOrigin struct {
	ProductID     string          `json:"productId"`
	CodeType      string          `json:"codeType"`
	User          User            `json:"user"`
	ProductPrices []ProductPrices `json:"productPrices"`
}

//User ...
type User struct {
	UserAccount string `json:"userAccount"`
	Mail        string `json:"mail"`
}

//ProductPrices ...
type ProductPrices struct {
	Stores    string      `json:"stores"`
	GscPrices []GscPrices `json:"prices"`
}

//GscPrices ...
type GscPrices struct {
	Type          int    `json:"type"`
	Value         string `json:"value"`
	Description   string `json:"description"`
	CurrencyCode  string `json:"currencyCode"`
	ProcessType   string `json:"processType"`
	StartDateTime string `json:"startDatetime"`
	EndDateTime   string `json:"endDatetime"`
}
