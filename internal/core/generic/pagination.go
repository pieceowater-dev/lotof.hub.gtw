package generic

import "app/internal/core/graph/model"

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

func IntToFriendshipStatus(stint int) model.FriendshipStatus {
	switch stint {
	case 100:
		return model.FriendshipStatusPending
	case 200:
		return model.FriendshipStatusAccepted
	case 300:
		return model.FriendshipStatusRejected
	default:
		return model.FriendshipStatusPending
	}
}

func FriendshipStatusToInt(st model.FriendshipStatus) int {
	switch st {
	case model.FriendshipStatusPending:
		return 100
	case model.FriendshipStatusAccepted:
		return 200
	case model.FriendshipStatusRejected:
		return 300
	default:
		return 100
	}
}
