package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// AutoSuggestFlights
//
// Returns a list of airports that match the given query.
//
// @see https://developers.skyscanner.net/docs/intro#geo-api
func AutoSuggestFlights(request requests.AutoSuggestRequest) (responses.PlacesResponse, error) {
	var response responses.PlacesResponse
	err := client.Json("POST", "autosuggest/flights", request, &response)
	return response, err
}
