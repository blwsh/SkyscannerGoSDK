package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// GetLocales
//
// Returns a list of locales supported by the API.
//
// @see https://developers.skyscanner.net/docs/culture/locales
func GetLocales() (responses.LocalesResponse, error) {
	var response responses.LocalesResponse
	err := client.Json("GET", "culture/locales", nil, &response)
	return response, err
}
