package flight

type Leg struct {
	OriginPlaceId       string   `json:"originPlaceId"`
	DestinationPlaceId  string   `json:"destinationPlaceId"`
	DepartureDateTime   DateTime `json:"departureDateTime"`
	ArrivalDateTime     DateTime `json:"arrivalDateTime"`
	DurationInMinutes   int      `json:"durationInMinutes"`
	StopCount           int      `json:"stopCount"`
	MarketingCarrierIds []string `json:"marketingCarrierIds"`
	OperatingCarrierIds []string `json:"operatingCarrierIds"`
	SegmentIds          []string `json:"segmentIds"`
}
