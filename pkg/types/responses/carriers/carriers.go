package carriers

import (
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types/flight"
)

type Carriers struct {
	Status   types.Status              `json:"status"`
	Carriers map[string]flight.Carrier `json:"carriers"`
}
