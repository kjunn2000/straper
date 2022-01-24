package board

type TaskBoardDataResponse struct {
	TaskBoard TaskBoard  `json:"task_board"`
	TaskLists []TaskList `json:"task_lists"`
}

type UpdateListParams struct {
	ListId   string `json:"list_id"`
	ListName string `json:"list_name"`
}

type UpdateCardParams struct {
	ListId      string `json:"list_id"`
	CardId      string `json:"card_id" db:"card_id"`
	Title       string `json:"title" db:"title"`
	Priority    string `json:"priority" db:"priority"`
	Description string `json:"description" db:"description"`
}

type DeleteCardParams struct {
	ListId string `json:"list_id"`
	CardId string `json:"card_id"`
}

type OrderListParams struct {
	BoardId      string `json:"board_id"`
	OldListIndex int    `json:"oldListIndex"`
	NewListIndex int    `json:"newListIndex"`
}

type OrderCardParams struct {
	SourceListId string `json:"sourceListId"`
	DestListId   string `json:"destListId"`
	OldCardIndex int    `json:"oldCardIndex"`
	NewCardIndex int    `json:"newCardIndex"`
}
