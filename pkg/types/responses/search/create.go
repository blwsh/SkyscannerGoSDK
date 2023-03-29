package search

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight/stats"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/sorting"
)

type CreateResponse struct {
	SessionToken string       `json:"sessionToken,omitempty"`
	Status       types.Status `json:"status"`
	Action       string       `json:"action"`
	Content      struct {
		Results struct {
			Itineraries map[string]flight.Itinerary `json:"itineraries,omitempty"`
			Legs        map[string]flight.Leg       `json:"legs,omitempty"`
			Segments    map[string]flight.Segment   `json:"segments,omitempty"`
			Places      map[string]geo.Place        `json:"places,omitempty"`
			Carriers    map[string]flight.Carrier   `json:"carriers,omitempty"`
			Agents      map[string]flight.Agent     `json:"agents,omitempty"`
			Alliances   map[string]flight.Alliance  `json:"alliances,omitempty"`
			Quotes      map[string]flight.Quote     `json:"quotes,omitempty"`
		}
	} `json:"content,omitempty"`
	Stats          stats.ItineraryStats `json:"stats,omitempty"`
	SortingOptions struct {
		Best     []sorting.Score `json:"best,omitempty"`
		Cheapest []sorting.Score `json:"cheapest,omitempty"`
		Fastest  []sorting.Score `json:"fastest,omitempty"`
	} `json:"sortingOptions,omitempty"`
}
