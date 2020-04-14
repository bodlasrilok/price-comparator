package vasproduct

type FilterParams struct {
	Offset          int
	Limit           int
	Keyword         string
	CategoryFilters []CategoryFilter
	SortFilters     []SortFilter
}

type CategoryFilter struct {
	Level int
	Value int64
}

type SortFilter struct {
	Field     string
	Ascending bool
}
