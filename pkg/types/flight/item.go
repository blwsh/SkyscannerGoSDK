package flight

type Item struct {
	Price    Price  `json:"price"`
	AgentId  string `json:"agentId"`
	DeepLink string `json:"deepLink"`
	Fares    []Fare `json:"fares"`
}
