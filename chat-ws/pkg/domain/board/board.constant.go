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

	BoardAddCard             = "BOARD_ADD_CARD"
	BoardUpdateCard          = "BOARD_UPDATE_CARD"
	BoardUpdateCardDueDate   = "BOARD_UPDATE_CARD_DUE_DATE"
	BoardUpdateCardIssueLink = "BOARD_UPDATE_CARD_ISSUE_LINK"
	BoardDeleteCard          = "BOARD_DELETE_CARD"
	BoardOrderCard           = "BOARD_ORDER_CARD"

	BoardCardAddMembers   = "BOARD_CARD_ADD_MEMBERS"
	BoardCardRemoveMember = "BOARD_CARD_REMOVE_MEMBER"

	BoardCardAddChecklistItem    = "BOARD_CARD_ADD_CHECKLIST_ITEM"
	BoardCardUpdateChecklistItem = "BOARD_CARD_UPDATE_CHECKLIST_ITEM"
	BoardCardDeleteChecklistItem = "BOARD_CARD_DELETE_CHECKLIST_ITEM"

	BoardCardAddComment    = "BOARD_CARD_ADD_COMMENT"
	BoardCardEditComment   = "BOARD_CARD_EDIT_COMMENT"
	BoardCardDeleteComment = "BOARD_CARD_DELETE_COMMENT"
)

const (
	TypeMessage = "MESSAGE"
	TypeFile    = "FILE"
)
