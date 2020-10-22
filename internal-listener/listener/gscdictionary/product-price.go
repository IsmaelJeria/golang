package gscdictionary

import (
	"encoding/json"
	"strconv"
	"strings"
)

//ProductPrice ...
type ProductPrice struct {
	Type           string   `json:"type"`
	StoreID        int      `json:"storeId"`
	SKU            int      `json:"sku"`
	Currency       string   `json:"currency"`
	CurrencyCode   string   `json:"currencyCode"`
	CurrencySymbol string   `json:"currencySymbol"`
	Country        string   `json:"country"`
	Prices         []Prices `json:"prices"`
}

//Prices ..
type Prices struct {
	Type          string `json:"type"`
	Value         string `json:"value"`
	StartDatetime string `json:"startDatetime"`
	EndDatetime   string `json:"endDatetime"`
	Description   string `json:"description"`
	ProcessType   string `json:"processType"`
}

//GscPOSTToProductPrice ...
func (d *ProductPrice) GscPOSTToProductPrice(data []byte) []byte {
	var gscPostOrigin []GscPostOrigin
	json.Unmarshal(data, &gscPostOrigin)

	var destinyList []ProductPrice

	for _, offerData := range gscPostOrigin[0].OfferData {
		var destiny ProductPrice
		destiny.Type = "PRICE" // fijo
		destiny.StoreID = 2000 // fijo
		sku, _ := strconv.Atoi(offerData.OfferingID)
		destiny.SKU = sku            // OfferingID
		destiny.Currency = "id"      // fijo
		destiny.CurrencyCode = "CLP" // fijo
		destiny.CurrencySymbol = "$" // fijo
		destiny.Country = "CL"       // fijo
		for _, pricedata := range offerData.PriceData {
			var offertype string
			if pricedata.PriceType == "OFFER" {
				offertype = "2"
			}
			if pricedata.PriceType == "NORMAL" {
				offertype = "1"
			}
			destiny.Prices = append(destiny.Prices, Prices{
				Type:          offertype,                                        // fijo
				Value:         strconv.Itoa(pricedata.Value),                    // priceData value
				StartDatetime: strings.ReplaceAll(pricedata.StartDate, "Z", ""), // fijo
				EndDatetime:   strings.ReplaceAll(pricedata.EndDate, "Z", ""),   // fijo
			})
		}
		destinyList = append(destinyList, destiny)
	}
	res, _ := json.Marshal(destinyList)
	return res
}

//GscPATCHToProductPrice ...
func (d *ProductPrice) GscPATCHToProductPrice(data []byte) []byte {

	var gscPatchOrigin GSCPatchOrigin
	json.Unmarshal(data, &gscPatchOrigin)

	var destinyList []ProductPrice

	var destiny ProductPrice
	destiny.Type = "PRICE" // fijo
	var stores []string
	destiny.StoreID = 0
	sku, _ := strconv.Atoi(gscPatchOrigin.ProductID)
	destiny.SKU = sku            //productId
	destiny.Currency = "id"      // fijo
	destiny.CurrencyCode = "CLP" // fijo
	destiny.CurrencySymbol = "$" // fijo
	destiny.Country = "CL"       // fijo
	for _, pricedata := range gscPatchOrigin.ProductPrices[0].GscPrices {
		destiny.Prices = append(destiny.Prices, Prices{
			Type:          strconv.Itoa(pricedata.Type),                         // fijo
			Value:         pricedata.Value,                                      // priceData value
			StartDatetime: strings.ReplaceAll(pricedata.StartDateTime, "Z", ""), // fijo
			EndDatetime:   strings.ReplaceAll(pricedata.EndDateTime, "Z", ""),   // fijo
			Description:   pricedata.Description,
			ProcessType:   pricedata.ProcessType,
		})
	}
	stores = strings.Split(gscPatchOrigin.ProductPrices[0].Stores, ", ")
	for _, store := range stores {
		s, _ := strconv.Atoi(store)
		destiny.StoreID = s
		destinyList = append(destinyList, destiny)
	}
	res, _ := json.Marshal(destinyList)
	return res
}
