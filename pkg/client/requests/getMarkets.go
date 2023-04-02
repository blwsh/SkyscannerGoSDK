package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// GetMarkets
//
// Returns a list of markets that Skyscanner supports.
//
// @see https://developers.skyscanner.net/docs/culture/markets
func GetMarkets() (responses.MarketsResponse, error) {
	var response responses.MarketsResponse
	err := client.Json("GET", "culture/locales", nil, &response)
	return response, err
}
