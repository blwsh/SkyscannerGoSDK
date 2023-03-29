package flight

type Fare struct {
	SegmentId     string `json:"segmentId"`
	BookingCode   string `json:"bookingCode"`
	FareBasisCode string `json:"fareBasisCode"`
}
