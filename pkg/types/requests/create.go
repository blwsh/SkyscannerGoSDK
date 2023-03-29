package requests

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"
)

type PlaceQuery struct {
	QueryPlace geo.Place `json:"queryPlace"`
}

type PlaceIdQuery struct {
	Iata string `json:"iata"`
}

type LegQuery struct {
	OriginPlaceId      PlaceIdQuery     `json:"originPlaceId"`
	DestinationPlaceId PlaceIdQuery     `json:"destinationPlaceId"`
	OriginPlace        PlaceQuery       `json:"originPlace"`
	DestinationPlace   PlaceQuery       `json:"destinationPlace"`
	Date               flight.Date      `json:"date"`
	DateRange          flight.DateRange `json:"date_range"`
}

type CreateQuery struct {
	Adults               int                    `json:"adults"`
	Market               string                 `json:"market"`
	Locale               string                 `json:"locale"`
	Currency             string                 `json:"currency"`
	CabinClass           string                 `json:"cabinClass"`
	QueryLegs            []LegQuery             `json:"queryLegs"`
	DateTimeGroupingType types.DateTimeGrouping `json:"dateTimeGroupingType"`
	ChildrenAges         []interface{}          `json:"childrenAges"`
	ExcludedAgentsIds    []interface{}          `json:"excludedAgentsIds"`
	ExcludedCarriersIds  []interface{}          `json:"excludedCarriersIds"`
	IncludedAgentsIds    []interface{}          `json:"includedAgentsIds"`
	IncludedCarriersIds  []interface{}          `json:"includedCarriersIds"`
}

type CreateQueryRequest struct {
	Query CreateQuery `json:"query"`
}
