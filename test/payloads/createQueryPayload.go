package payloads

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	searchQuery "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
)

var CreateQueryPayload = searchQuery.CreateQueryRequest{
	Query: searchQuery.CreateQuery{
		Adults:               1,
		Market:               "US",
		Currency:             "USD",
		Locale:               "en-US",
		DateTimeGroupingType: types.ByMonth,
		CabinClass:           "CABIN_CLASS_ECONOMY",
		QueryLegs: []searchQuery.LegQuery{
			{
				OriginPlaceId:      searchQuery.PlaceIdQuery{Iata: "LHR"},
				DestinationPlaceId: searchQuery.PlaceIdQuery{Iata: "EDI"},
				Date: flight.Date{
					Year:  2023,
					Month: 12,
					Day:   01,
				},
			},
		},
	},
}
