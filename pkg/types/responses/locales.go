package responses

import "github.com/blwsh/SkyscannerGoSDK/pkg/types"

type LocalesResponse struct {
	types.Status
	Locales []types.Locale `json:"locales"`
}
