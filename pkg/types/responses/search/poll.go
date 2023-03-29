package search

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight/stats"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/sorting"
)

type PollResponse struct {
	SessionToken string       `json:"session_token,omitempty"`
	Status       types.Status `json:"status"`
	Action       string       `json:"action"`
	Content      struct {
		Results struct {
			Itineraries map[string]flight.Itinerary
			Legs        map[string]flight.Leg
			Segments    map[string]flight.Segment
			Places      map[string]geo.Place
			Carriers    map[string]flight.Carrier
			Agents      map[string]flight.Agent
			Alliances   map[string]flight.Alliance
		}
	}
	Stats          stats.ItineraryStats
	SortingOptions struct {
		Best     []sorting.Score
		Cheapest []sorting.Score
		Fastest  []sorting.Score
	}
}
