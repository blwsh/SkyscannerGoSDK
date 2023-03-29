package culture

import "github.com/blwsh/SkyscannerGoSDK/pkg/types"

type Market struct {
	Status  types.Status `json:"status"`
	Markets []Market     `json:"markets"`
}
