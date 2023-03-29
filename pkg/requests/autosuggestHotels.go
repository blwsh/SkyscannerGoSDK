package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/internal/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// AutoSuggestHotels
//
// Returns a list of hotels locations that match the given query.
//
// @see https://developers.skyscanner.net/docs/intro#geo-api
func AutoSuggestHotels(request requests.AutoSuggestRequest) (responses.PlacesResponse, error) {
	var response responses.PlacesResponse
	err := client.Json("POST", "autosuggest/carhire", request, &response)
	return response, err
}
