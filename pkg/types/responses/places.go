package responses

import "github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"

type PlacesResponse struct {
	Places []geo.Place `json:"places"`
}
