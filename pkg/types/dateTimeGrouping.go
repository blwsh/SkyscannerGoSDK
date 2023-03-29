package types

type DateTimeGrouping string

const (
	// Unspecified The date grouping is not specified (use the default).
	Unspecified = "DATE_TIME_GROUPING_TYPE_UNSPECIFIED"
	// ByDate Group the quotes by day.
	ByDate = "DATE_TIME_GROUPING_TYPE_BY_DATE"
	// ByMonth Group the quotes by month.
	ByMonth = "DATE_TIME_GROUPING_TYPE_BY_MONTH"
)
