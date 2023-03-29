package indicativePrices

type QuoteGroup struct {
	OriginPlaceId      string   `json:"originPlaceId"`
	DestinationPlaceId string   `json:"destinationPlaceId"`
	QuoteIds           []string `json:"quoteIds"`
}
