package generic

import "app/internal/core/graph/model"

type SortDirection string

const (
	Asc  SortDirection = "ASC"
	Desc SortDirection = "DESC"
)

type Sort[T any] struct {
	Field     string        `json:"field"`
	Direction SortDirection `json:"direction"`
}

func NewSort[T any](field string, direction SortDirection) Sort[T] {
	return Sort[T]{
		Field:     field,
		Direction: direction,
	}
}

func SortByEnumToString(sortBy *model.FilterSortByEnum) string {
	if sortBy == nil {
		return "ASC"
	}
	switch *sortBy {
	case model.FilterSortByEnumAsc:
		return "ASC"
	case model.FilterSortByEnumDesc:
		return "DESC"
	default:
		return "ASC"
	}
}
