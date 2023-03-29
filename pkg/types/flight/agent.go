package flight

type Agent struct {
	Name                 string          `json:"name,omitempty"`
	Type                 string          `json:"type,omitempty"`
	ImageUrl             string          `json:"imageUrl,omitempty"`
	FeedbackCount        int             `json:"feedbackCount,omitempty"`
	Rating               float64         `json:"rating,omitempty"`
	RatingBreakdown      RatingBreakdown `json:"ratingBreakdown,omitempty"`
	IsOptimisedForMobile bool            `json:"isOptimisedForMobile,omitempty"`
}
