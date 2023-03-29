package culture

import "github.com/blwsh/SkyscannerGoSDK/pkg/types"

type Currencies struct {
	Status     types.Status     `json:"status"`
	Currencies []types.Currency `json:"currencies"`
}
