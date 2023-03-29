package flight

type Segment struct {
	OriginPlaceId         string   `json:"originPlaceId"`
	DestinationPlaceId    string   `json:"destinationPlaceId"`
	DepartureDateTime     DateTime `json:"departureDateTime"`
	ArrivalDateTime       DateTime `json:"arrivalDateTime"`
	DurationInMinutes     int      `json:"durationInMinutes"`
	MarketingFlightNumber string   `json:"marketingFlightNumber"`
	MarketingCarrierId    string   `json:"marketingCarrierId"`
	OperatingCarrierId    string   `json:"operatingCarrierId"`
}
