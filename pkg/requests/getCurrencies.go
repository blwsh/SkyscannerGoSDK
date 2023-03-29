package requests

import (
	client "github.com/blwsh/SkyscannerGoSDK/pkg/client"
	"github.com/blwsh/SkyscannerGoSDK/pkg/types"
)

type CurrenciesResponse struct {
	types.Status
	Currencies []types.Currency `json:"currencies"`
}

func GetCurrencies() (CurrenciesResponse, error) {
	var response CurrenciesResponse
	err := client.Json("GET", "culture/currencies", nil, &response)
	return response, err
}
