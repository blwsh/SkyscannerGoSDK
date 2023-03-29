package responses

import "github.com/blwsh/SkyscannerGoSDK/pkg/types"

type MarketsResponse struct {
	types.Status
	Markets []types.Market `json:"locales"`
}
