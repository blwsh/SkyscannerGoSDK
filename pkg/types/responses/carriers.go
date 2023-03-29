package responses

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
)

type CarriersResponse struct {
	Status   types.Status              `json:"status"`
	Carriers map[string]flight.Carrier `json:"carriers"`
}
