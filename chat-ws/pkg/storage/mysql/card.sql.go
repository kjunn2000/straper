package mysql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"go.uber.org/zap"
)

func (q *Queries) CreateCard(ctx context.Context, card board.Card) error {
	sql, args, err := sq.Insert("card").Columns("card_id", "title", "priority", "list_id",
		"description", "creator_id", "created_date", "due_date", "order_index").
		Values(card.CardId, card.Title, card.Priority, card.ListId, card.Description,
			card.CreatorId, card.CreatedDate, card.DueDate, card.OrderIndex).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert card sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to create new card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetCardListByListId(ctx context.Context, listId string) ([]board.Card, error) {
	sql, args, err := sq.Select("card_id", "title", "priority", "list_id", "description", "creator_id", "created_date",
		"due_date", "order_index").From("card").
		Where(sq.Eq{"list_id": listId}).
		OrderBy("order_index").
		ToSql()
	if err != nil {
		q.log.Info("Unable to create select card sql.", zap.Error(err))
		return []board.Card{}, err
	}
	var cardList []board.Card
	err = q.db.Select(&cardList, sql, args...)
	if err != nil {
		q.log.Info("Unable to create select card sql.", zap.Error(err))
		return []board.Card{}, err
	}
	return cardList, nil
}

func (q *Queries) UpdateCard(ctx context.Context, params board.UpdateCardParams) error {
	sql, args, err := sq.Update("card").
		Set("title", params.Title).
		Set("priority", params.Priority).
		Set("description", params.Description).
		Where(sq.Eq{"card_id": params.CardId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update card sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) UpdateCardDueDate(ctx context.Context, params board.UpdateCardDueDateParams) error {
	fmt.Println(params.DueDate)
	fmt.Println(params.CardId)
	sql, args, err := sq.Update("card").
		Set("due_date", params.DueDate).
		Where(sq.Eq{"card_id": params.CardId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update card sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) UpdateCardOrder(ctx context.Context, cardId string, orderIndex int, listId string, updateListId bool) error {
	updateBuilder := sq.Update("card").
		Set("order_index", orderIndex)
	if updateListId {
		updateBuilder = updateBuilder.Set("list_id", listId)
	}
	sql, args, err := updateBuilder.Where(sq.Eq{"card_id": cardId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update card order sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update card order.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteCard(ctx context.Context, cardId string) error {
	sql, args, err := sq.Delete("card").Where(sq.Eq{"card_id": cardId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create delete card sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetUserFromCard(ctx context.Context, cardId string) ([]string, error) {
	sql, args, err := sq.Select("user_id").From("card_user").
		Where(sq.Eq{"card_id": cardId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create select card user sql.", zap.Error(err))
		return []string{}, err
	}
	var userList []string
	err = q.db.Select(&userList, sql, args...)
	if err != nil {
		q.log.Info("Unable to create select card user sql.", zap.Error(err))
		return []string{}, err
	}
	return userList, nil
}

func (q *Queries) AddUserListToCard(ctx context.Context, cardId string, userIdList []string) error {
	builder := sq.Insert("card_user").Columns("card_id", "user_id")
	for _, userId := range userIdList {
		builder = builder.Values(cardId, userId)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		q.log.Info("Unable to create insert sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to add users to card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteUserFromCard(ctx context.Context, cardId, userId string) error {
	sql, args, err := sq.Delete("card_user").
		Where(sq.Eq{"card_id": cardId}).
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create delete sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete user from card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetChecklistItemsByCardId(ctx context.Context, cardId string) ([]string, error) {
	sql, args, err := sq.Select("item_id", "content", "is_checked", "card_id").From("checklist_item").
		Where(sq.Eq{"card_id": cardId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create select checklist items sql.", zap.Error(err))
		return []string{}, err
	}
	var userList []string
	err = q.db.Select(&userList, sql, args...)
	if err != nil {
		q.log.Info("Unable to create select checklist items sql.", zap.Error(err))
		return []string{}, err
	}
	return userList, nil
}

func (q *Queries) CreateChecklistItem(ctx context.Context, checklistItem board.CardChecklistItem) error {
	sql, args, err := sq.Insert("checklist_item").
		Columns("item_id", "content", "is_checked", "card_id").
		Values(checklistItem.CardId, checklistItem.Content, checklistItem.IsChecked, checklistItem.CardId).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create insert sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to add checklist item to card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) UpdateChecklistItem(ctx context.Context, checklistItem board.CardChecklistItem) error {
	sql, args, err := sq.Update("checklist_item").
		Set("content", checklistItem.Content).
		Set("is_checked", checklistItem.IsChecked).
		ToSql()
	if err != nil {
		q.log.Info("Failed to create update checklist sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update checklist.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteChecklistItem(ctx context.Context, itemId string) error {
	sql, args, err := sq.Delete("checklist_item").
		Where(sq.Eq{"item_id": itemId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create delete sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete checklist item from card.", zap.Error(err))
		return err
	}
	return nil
}
