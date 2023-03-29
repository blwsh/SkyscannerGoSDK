package flight

type RatingBreakdown struct {
	CustomerService float64 `json:"customerService"`
	ReliablePrices  float64 `json:"reliablePrices"`
	ClearExtraFees  float64 `json:"clearExtraFees"`
	EaseOfBooking   float64 `json:"easeOfBooking"`
	Other           float64 `json:"other"`
}
