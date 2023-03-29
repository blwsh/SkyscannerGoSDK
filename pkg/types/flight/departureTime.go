package flight

type Date struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

type DateTime struct {
	Date
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
	Second int `json:"second,omitempty"`
}

type DateRange struct {
	StartDate Date `json:"startDate"`
	EndDate   Date `json:"endDate"`
}
