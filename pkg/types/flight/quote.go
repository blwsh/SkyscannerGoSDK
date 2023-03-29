package flight

type Quote struct {
	MinPrice    Price `json:"minPrice"`
	IsDirect    bool  `json:"isDirect"`
	OutboundLeg Leg   `json:"outboundLeg"`
	InboundLeg  Leg   `json:"inboundLeg"`
}
