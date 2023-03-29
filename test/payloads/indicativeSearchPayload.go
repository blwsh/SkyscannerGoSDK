package payloads

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"
	searchQuery "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"time"
)

var IndicativeSearchPayload = searchQuery.CreateQueryRequest{
	Query: searchQuery.CreateQuery{
		Market:               "US",
		Currency:             "USD",
		Locale:               "en-US",
		DateTimeGroupingType: types.ByMonth,
		QueryLegs: []searchQuery.LegQuery{
			{
				OriginPlace:      searchQuery.PlaceQuery{QueryPlace: geo.Place{EntityId: "27544008"}},
				DestinationPlace: searchQuery.PlaceQuery{QueryPlace: geo.Place{EntityId: "27539733"}},
				DateRange: flight.DateRange{
					StartDate: flight.Date{
						Year:  time.Now().Year(),
						Month: int(time.Now().Month()),
					},
					EndDate: flight.Date{
						Year:  time.Now().Add(24 * time.Hour * 30).Year(),
						Month: int(time.Now().Add(24 * time.Hour * 30).Month()),
					},
				},
			},
		},
	},
}
