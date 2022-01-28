package board

import (
	"encoding/json"
	"time"
)

type TaskBoard struct {
	BoardId     string `json:"board_id" db:"board_id"`
	BoardName   string `json:"board_name" db:"board_name"`
	WorkspaceId string `json:"workspace_id" db:"workspace_id"`
}

type TaskList struct {
	ListId     string `json:"list_id" db:"list_id"`
	ListName   string `json:"list_name" db:"list_name"`
	BoardId    string `json:"board_id" db:"board_id"`
	OrderIndex int    `json:"order_index" db:"order_index"`
	CardList   []Card `json:"card_list"`
}

type Card struct {
	CardId      string              `json:"card_id" db:"card_id"`
	Title       string              `json:"title" db:"title"`
	Priority    string              `json:"priority" db:"priority"`
	ListId      string              `json:"list_id" db:"list_id"`
	Description string              `json:"description" db:"description"`
	CreatorId   string              `json:"creator_id" db:"creator_id"`
	CreatedDate time.Time           `json:"created_date" db:"created_date"`
	DueDate     time.Time           `json:"due_date" db:"due_date"`
	OrderIndex  int                 `json:"order_index" db:"order_index"`
	MemberList  []string            `json:"member_list"`
	Checklist   []CardChecklistItem `json:"checklist"`
}

func (card *Card) Encode() ([]byte, error) {
	json, err := json.Marshal(card)
	if err != nil {
		return nil, err
	}
	return json, nil
}

type CardChecklistItem struct {
	ItemId    string `json:"item_id" db:"item_id"`
	Content   string `json:"content" db:"content"`
	IsChecked bool   `json:"is_checked" db:"is_checked"`
	CardId    string `json:"card_id" db:"card_id"`
}

type CardComment struct {
	CommentId   string     `json:"comment_id" db:"comment_id"`
	Type        string     `json:"type" db:"type"`
	CardId      string     `json:"card_id" db:"card_id"`
	CreatorId   string     `json:"creator_id" db:"creator_id"`
	UserDetail  UserDetail `json:"user_detail"`
	Content     string     `json:"content" db:"content"`
	FileName    string     `json:"file_name" db:"file_name"`
	FileType    string     `json:"file_type" db:"file_type"`
	FileBytes   []byte     `json:"file_bytes"`
	CreatedDate time.Time  `json:"created_date" db:"created_date"`
}

type UserDetail struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email" validate:"email"`
	PhoneNo  string `json:"phone_no" db:"phone_no"`
}
