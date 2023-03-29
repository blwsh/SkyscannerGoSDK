package requests

import (
	"fmt"
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// GetFlightsHierarchy
//
// The Geo API returns a full catalogue of supported locations (airports, cities, islands, countries and continents)
// in the specified locale. Supported locations are locations where travelers can search flights to and from.
//
// @see https://developers.skyscanner.net/docs/geo/quick-start#hierarchyflights
func GetFlightsHierarchy() (responses.FlightsHierarchyResponse, error) {
	var response responses.FlightsHierarchyResponse
	err := client.Json("GET", fmt.Sprintf("geo/hierarchy/flights/%s", client.GetLocale()), nil, &response)
	return response, err
}
