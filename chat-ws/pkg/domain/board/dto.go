package board

import "time"

type TaskBoardDataResponse struct {
	TaskBoard TaskBoard  `json:"task_board"`
	TaskLists []TaskList `json:"task_lists"`
}

type UpdateCardParams struct {
	CardId      string    `json:"card_id" db:"card_id"`
	Title       string    `json:"title" db:"title"`
	Status      string    `json:"status" db:"status"`
	Priority    string    `json:"priority" db:"priority"`
	ListId      string    `json:"list_id"`
	Description string    `json:"description" db:"description"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
}

type UpdateCardOrderParams struct {
	CardId     string `json:"card_id" db:"card_id"`
	OrderIndex int    `json:"order_index" db:"order_index"`
}
