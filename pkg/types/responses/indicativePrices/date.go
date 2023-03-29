package indicativePrices

type Date struct {
	QuotesInboundGroups  []interface{} `json:"quotesInboundGroups"`
	QuotesOutboundGroups []struct {
		MonthYearDate struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"monthYearDate"`
		QuoteIds []string `json:"quoteIds"`
	} `json:"quotesOutboundGroups"`
}
