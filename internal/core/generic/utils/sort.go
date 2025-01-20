package utils

import "app/internal/core/graph/model"

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
