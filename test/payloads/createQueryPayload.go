package payloads

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
)

var CreateQueryPayload = requests.CreateQueryRequest{
	Query: requests.CreateQuery{
		Adults:               1,
		Market:               "US",
		Currency:             "USD",
		Locale:               "en-US",
		DateTimeGroupingType: types.ByMonth,
		CabinClass:           "CABIN_CLASS_ECONOMY",
		QueryLegs: []requests.LegQuery{
			{
				OriginPlaceId:      requests.PlaceIdQuery{Iata: "MAN"},
				DestinationPlaceId: requests.PlaceIdQuery{Iata: "HAM"},
				Date: flight.Date{
					Year:  2023,
					Month: 12,
					Day:   01,
				},
			},
		},
	},
}
