package types

type ResultStats string

const (
	// ResultStatusUnspecified Unspecified Results status not specified.
	ResultStatusUnspecified = "RESULT_STATUS_UNSPECIFIED"
	// ResultStatusComplete Complete Results are now complete.
	ResultStatusComplete = "RESULT_STATUS_COMPLETE"
	// ResultStatusIncomplete Incomplete Results are not complete yet.
	ResultStatusIncomplete = "RESULT_STATUS_INCOMPLETE"
	// ResultStatusFailed Failed Search has failed.
	ResultStatusFailed = "RESULT_STATUS_FAILED"
)
