package geo

type Place struct {
	EntityId     string  `json:"entityId"`
	IataCode     string  `json:"iataCode"`
	ParentId     string  `json:"parentId"`
	Name         string  `json:"name"`
	CountryId    string  `json:"countryId"`
	CountryName  string  `json:"countryName"`
	CityName     string  `json:"cityName"`
	Location     string  `json:"location"`
	Hierarchy    string  `json:"hierarchy"`
	Type         string  `json:"type"`
	Highlighting [][]int `json:"highlighting"`
}
