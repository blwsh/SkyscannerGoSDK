package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	searchQuery "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses/search"
)

// IndicativeFlightSearch
//
// The Indicative Prices API returns a list of the cheapest prices seen last by our travellers for a given search criteria.
//
// The Flights Indicative Prices API allows for more flexibility than the Flights Live Prices API. You can access the
// quotes based on the aggregation categories we support.
//
// We support two aggregation categories: date aggregations and
// route aggregations. Each of them have some sub-categories.
//
// @see https://developers.skyscanner.net/docs/flights-indicative-prices/overview
func IndicativeFlightSearch(request searchQuery.CreateQueryRequest) (search.CreateResponse, error) {
	var createResponse search.CreateResponse
	err := client.Json("POST", "flights/indicative/search", request, &createResponse)
	return createResponse, err
}
