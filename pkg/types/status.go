package types

// Status
//
// The response that is returned is different depending on the endpoint but the status is used across all endpoints.
// This param lets you know the status of the results received if its complete or its partial data of a search.
//
// @see https://developers.skyscanner.net/docs/getting-started/requests-and-responses
type Status string

const (
	// Incomplete The search has not been yet completed. Used for search's that can take time, there will be a /poll endpoint to check on the search status and get the extra data.
	Incomplete Status = "RESULT_STATUS_INCOMPLETE"
	// Complete All data have been retrieved.
	Complete Status = "RESULT_STATUS_COMPLETE"
	// Failed An error has occurred while retrieving data.
	Failed Status = "RESULT_STATUS_FAILED"
)

func (s Status) IsValid() bool {
	return s == Incomplete || s == Complete || s == Failed
}

func (s Status) IsFinal() bool {
	return s == Complete || s == Failed
}

func (s Status) IsComplete() bool {
	return s == Complete
}

func (s Status) IsFailed() bool {
	return s == Failed
}

func (s Status) IsInomplete() bool {
	return s == Incomplete
}
