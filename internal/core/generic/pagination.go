package generic

import "app/internal/core/graph/model"

type Pagination struct {
	Page   int `json:"page"`
	Length int `json:"length"`
}

func NewPagination(page, length int) Pagination {
	return Pagination{
		Page:   page,
		Length: length,
	}
}

func PaginationLengthToInt(length model.FilterPaginationLengthEnum) int32 {
	switch length {
	case model.FilterPaginationLengthEnumTen:
		return 10
	case model.FilterPaginationLengthEnumFifteen:
		return 15
	case model.FilterPaginationLengthEnumTwenty:
		return 20
	case model.FilterPaginationLengthEnumTwentyFive:
		return 25
	case model.FilterPaginationLengthEnumThirty:
		return 30
	case model.FilterPaginationLengthEnumThirtyFive:
		return 35
	case model.FilterPaginationLengthEnumForty:
		return 40
	case model.FilterPaginationLengthEnumFortyFive:
		return 45
	case model.FilterPaginationLengthEnumFifty:
		return 50
	case model.FilterPaginationLengthEnumFiftyFive:
		return 55
	case model.FilterPaginationLengthEnumSixty:
		return 60
	case model.FilterPaginationLengthEnumSixtyFive:
		return 65
	case model.FilterPaginationLengthEnumSeventy:
		return 70
	case model.FilterPaginationLengthEnumSeventyFive:
		return 75
	case model.FilterPaginationLengthEnumEighty:
		return 80
	case model.FilterPaginationLengthEnumEightyFive:
		return 85
	case model.FilterPaginationLengthEnumNinety:
		return 90
	case model.FilterPaginationLengthEnumNinetyFive:
		return 95
	case model.FilterPaginationLengthEnumOneHundred:
		return 100
	default:
		return 10
	}
}
