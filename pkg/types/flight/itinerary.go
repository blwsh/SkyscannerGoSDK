package flight

type Itinerary struct {
	PricingOptions []struct {
		Price struct {
			Amount       string `json:"amount"`
			Unit         string `json:"unit"`
			UpdateStatus string `json:"updateStatus"`
		} `json:"price"`
		AgentIds     []string `json:"agentIds"`
		Items        []Item   `json:"items"`
		TransferType string   `json:"transferType"`
		Id           string   `json:"id"`
	} `json:"pricingOptions"`
	LegIds             []string    `json:"legIds"`
	SustainabilityData interface{} `json:"sustainabilityData"`
}
