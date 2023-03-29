package requests

type AutoSuggestRequest struct {
	Query AutoSuggestQuery `json:"query"`
	Limit int              `json:"limit"`
}

type AutoSuggestQuery struct {
	Market     string `json:"market"`
	Locale     string `json:"locale"`
	SearchTerm string `json:"searchTerm"`
}
