package indicativePrices

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"
)

type Search struct {
	Status  types.Status `json:"status"`
	Content struct {
		Results struct {
			Quotes   map[string][]flight.Quote   `json:"quotes"`
			Carriers map[string][]flight.Carrier `json:"carriers"`
			Places   map[string][]geo.Place      `json:"places"`
		} `json:"results"`
		GroupingOptions struct {
			ByRoute struct {
				QuotesGroup []QuoteGroup `json:"quotes_group"`
			} `json:"byRoute"`
			ByDate Date `json:"byDate"`
		}
	}
}
