package responses

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/geo"
)

type FlightsHierarchyResponse struct {
	Status types.Status         `json:"status"`
	Places map[string]geo.Place `json:"places"`
}
