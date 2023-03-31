package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/requests"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// AutoSuggestCarhire
//
// Returns a list of car hire locations that match the given query.
//
// @see https://developers.skyscanner.net/docs/intro#geo-api
func AutoSuggestCarhire(request requests.AutoSuggestRequest) (responses.PlacesResponse, error) {
	var response responses.PlacesResponse
	err := client.Json("POST", "autosuggest/carhire", request, &response)
	return response, err
}
