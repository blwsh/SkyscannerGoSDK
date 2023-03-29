package flight

type Carrier struct {
	Name       string `json:"name"`
	AllianceId string `json:"allianceId"`
	ImageUrl   string `json:"imageUrl"`
	Iata       string `json:"iata"`
}
