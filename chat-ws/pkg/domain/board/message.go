package board

import (
	"encoding/json"
	"time"
)

type Card struct {
	CardId      string    `json:"card_id" db:"card_id"`
	Title       string    `json:"title" db:"title"`
	Status      string    `json:"status" db:"status"`
	Priority    string    `json:"priority" db:"priority"`
	ListId      string    `json:"list_id"`
	Description string    `json:"description" db:"description"`
	CreatorId   string    `json:"creator_id" db:"creator_id"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
}

func (card *Card) Encode() ([]byte, error) {
	json, err := json.Marshal(card)
	if err != nil {
		return nil, err
	}
	return json, nil
}
