package types

type Currency struct {
	Code                        string `json:"code"`
	Symbol                      string `json:"symbol"`
	ThousandsSeparator          string `json:"thousandsSeparator"`
	DecimalSeparator            string `json:"decimalSeparator"`
	SymbolOnLeft                bool   `json:"symbolOnLeft"`
	SpaceBetweenAmountAndSymbol bool   `json:"spaceBetweenAmountAndSymbol"`
	DecimalDigits               int    `json:"decimalDigits"`
}
