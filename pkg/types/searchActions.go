package types

// SearchAction
//
// The action that has been performed on the search.
//
// @see https://developers.skyscanner.net/docs/getting-started/enums#searchaction
type SearchAction string

const (
	// SearchActionUnspecified Unspecified The action is not specified.
	SearchActionUnspecified = "SEARCH_ACTION_UNSPECIFIED"
	// SearchActionCreate The search has been created.
	SearchActionCreate = "SEARCH_ACTION_CREATE"
	// SearchActionPoll The search has been updated.
	SearchActionPoll = "SEARCH_ACTION_POLL"
)
