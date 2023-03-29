package culture

import "github.com/blwsh/SkyscannerGoSDK/pkg/types"

type Locales struct {
	Status  types.Status   `json:"status"`
	Locales []types.Locale `json:"locales"`
}
