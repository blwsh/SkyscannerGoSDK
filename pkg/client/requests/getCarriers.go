package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/responses"
)

// GetCarriers
//
// The Carriers API returns a full list of active carriers with name and IATA code indexed by their carrierId.
//
// This is the same carrierId from Flights Live Prices API reference and Flights Indicative API reference. The carrierId
// is in the same format e.g. -32677.
//
// This endpoint can be used to get up-to-date IATA code mappings following the internal Skyscanner carriers database.
//
// @see https://developers.skyscanner.net/docs/carriers/overview
func GetCarriers() (responses.CarriersResponse, error) {
	var carriersResponse responses.CarriersResponse
	err := client.Json("GET", "flights/carriers", nil, &carriersResponse)
	return carriersResponse, err
}
