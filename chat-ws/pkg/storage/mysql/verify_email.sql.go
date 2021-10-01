package mysql

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"go.uber.org/zap"
)

func (q *Queries) CreateVerifyEmailToken(ctx context.Context, token account.VerifyEmailToken) error {
	sql, args, err := sq.Insert("verify_email_token").Columns("token_id", "user_id", "created_date").Values(token.TokenId, token.UserId, time.Now()).ToSql()
	if err != nil {
		q.log.Warn("Fail to create verify email token query", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Warn("Fail to create verify email token", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetVerifyEmailToken(ctx context.Context, tokenId string) (account.VerifyEmailToken, error) {
	sql, args, err := sq.Select("token_id", "user_id", "created_date").From("verify_email_token").Where(sq.Eq{"token_id": tokenId}).ToSql()
	if err != nil {
		q.log.Warn("Fail to get verify email token query", zap.Error(err))
		return account.VerifyEmailToken{}, err
	}
	var token account.VerifyEmailToken
	err = q.db.Get(&token, sql, args...)
	if err != nil {
		q.log.Warn("Fail to get verify email token", zap.Error(err))
		return account.VerifyEmailToken{}, err
	}
	return token, nil
}

func (q *Queries) DeleteVerifyEmailToken(ctx context.Context, tokenId string) error {
	sql, args, err := sq.Delete("verify_email_token").Where(sq.Eq{"token_id": tokenId}).ToSql()
	if err != nil {
		q.log.Warn("Fail to create get verify email token query.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Warn("Fail to get verify email token query", zap.Error(err))
		return err
	}
	return nil
}
