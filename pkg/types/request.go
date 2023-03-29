package types

// Request
//
// The APIs use URL params or a query object which will be sent via a POST request. The query object allows you to define
// the search like currency, market, language, dates, one way, etc. This object changes based on the specific endpoint
// so please refer to the documentation of the query object for each endpoint you are using.
//
// The query object has some common params such as market, locale and currency, this data can also be received from the
// Culture API which you will learn about in the next section. [See Culture information](https://developers.skyscanner.net/docs/getting-started/culture-service)
//
// @see https://developers.skyscanner.net/docs/getting-started/requests-and-responses
type Request struct {
	// The market you are interested in. The market is typically the country where the user is making the search from.
	// For example, if a user in the UK is searching for a flight to Spain, the market would be "UK" and the currency "GBP".
	// The market must be a supported market, which you can find using the List Supported Markets API.
	Market string `json:"market"`
	// The locale you would like the results in. The locale must be a supported locale, which you can find using the
	// List Supported Locales API.
	Locale string `json:"locale"`
	// The currency you would like the prices in. The currency must be a supported currency, which you can find using the
	// List Supported Currencies API.
	Currency string `json:"currency"`
}

// SessionRequest
//
// The session request is used to create a session for a search. The session is used to poll for the results of the search.
// The session is created by sending a POST request to the /sessions endpoint with the query object as the body.
//
// @see https://developers.skyscanner.net/docs/getting-started/requests-and-responses
type SessionRequest struct {
	Request `json:",inline"`
}
