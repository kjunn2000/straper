package board

const (
	NoPriority     = "NO"
	LowPriority    = "LOW"
	MediumPriority = "MEDIUM"
	HighPriority   = "HIGH"
)

const (
	BoardAddList    = "BOARD_ADD_LIST"
	BoardUpdateList = "BOARD_UPDATE_LIST"
	BoardDeleteList = "BOARD_DELETE_LIST"
	BoardOrderList  = "BOARD_ORDER_LIST"

	BoardAddCard           = "BOARD_ADD_CARD"
	BoardUpdateCard        = "BOARD_UPDATE_CARD"
	BoardUpdateCardDueDate = "BOARD_UPDATE_CARD_DUE_DATE"
	BoardDeleteCard        = "BOARD_DELETE_CARD"
	BoardOrderCard         = "BOARD_ORDER_CARD"
	BoardCardAddMembers    = "BOARD_CARD_ADD_MEMBERS"
	BoardCardRemoveMember  = "BOARD_CARD_REMOVE_MEMBER"
)
