package gscdictionary

//GscPostOrigin ...
type GscPostOrigin struct {
	VariantID string      `json:"variantId"`
	OfferData []OfferData `json:"offerData"`
}

//OfferData ...
type OfferData struct {
	OfferingID string      `json:"offeringId"`
	PriceData  []PriceData `json:"priceData"`
}

//PriceData ...
type PriceData struct {
	PriceType       string       `json:"priceType"`
	PriceGroup      string       `json:"priceGroup"`
	Value           int          `json:"value"`
	CountryCode     string       `json:"countryCode"`
	CurrencyCode    string       `json:"currencyCode"`
	SourceUpdatedAt string       `json:"sourceUpdatedAt"`
	Precision       int          `json:"precision"`
	StartDate       string       `json:"startDate"`
	EndDate         string       `json:"endDate"`
	Priority        int          `json:"priority"`
	Attributes      []Attributes `json:"attributes"`
	IsPublished     bool         `json:"isPublished"`
}

//Attributes ...
type Attributes struct {
	Name   string   `json:"name"`
	Values []Values `json:"values"`
}

//Values ...
type Values struct {
	String string `json:"string"`
}
