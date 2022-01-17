package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"go.uber.org/zap"
)

func (q *Queries) CreateCard(ctx context.Context, card board.Card) error {
	sql, args, err := sq.Insert("card").Columns("card_id", "title", "status", "priority", "list_id",
		"description", "creator_id", "created_date", "due_data", "order").
		Values(card.CardId, card.Title, card.Status, card.Priority, card.ListId, card.Description,
			card.CreatorId, card.CreatedDate, card.DueDate, card.Order).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert card sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to create new card list.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetCardListByListId(ctx context.Context, listId string) ([]board.Card, error) {
	sql, args, err := sq.Select("card_id", "title", "status", "priority", "list_id",
		"description", "creator_id", "created_date", "due_data", "order").From("card").
		Where(sq.Eq{"list_id": listId}).ToSql()
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
		Set("status", params.Status).
		Set("priority", params.Priority).
		Set("list_id", params.ListId).
		Set("description", params.Description).
		Set("due_date", params.DueDate).
		Set("order", params.Order).
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

func (q *Queries) AddUserToCard(ctx context.Context, cardId, userId string) error {
	sql, args, err := sq.Insert("card_user").Columns("card_id", "user_id").
		Values(cardId, userId).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to user to card.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteUserFromCard(ctx context.Context, cardId, userId string) error {
	sql, args, err := sq.Delete("card").
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
