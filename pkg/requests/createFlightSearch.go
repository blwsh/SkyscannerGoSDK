package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	searchQuery "github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses/search"
)

// CreateFlightSearch
//
// Used to initiate the search request. This endpoint returns an incomplete cached subset of results for a quicker time
// to first result.
//
// @see https://developers.skyscanner.net/docs/flights-live-prices/overview#endpoints
func CreateFlightSearch(request searchQuery.CreateQueryRequest) (search.CreateResponse, error) {
	var createResponse search.CreateResponse
	err := client.Json("POST", "flights/live/search/create", request, &createResponse)
	return createResponse, err
}
