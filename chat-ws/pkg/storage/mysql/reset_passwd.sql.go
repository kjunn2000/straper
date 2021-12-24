package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"go.uber.org/zap"
)

func (q *Queries) CreateResetPasswordToken(ctx context.Context, params account.ResetPasswordToken) error {
	sql, arg, err := sq.Insert("reset_password_token").
		Columns("tokenId", "userId", "created_date").
		Values(params.TokenId, params.UserId, params.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Failed to create reset passwd token query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Info("Failed to create passwd token record to db.", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (q *Queries) GetResetPasswordToken(ctx context.Context, tokenId string) (account.ResetPasswordToken, error) {
	sql, args, err := sq.Select("token_id", "user_id", "created_date").From("reset_password_token").Where(sq.Eq{"token_id": tokenId}).ToSql()
	if err != nil {
		q.log.Warn("Fail to get reset password token query", zap.Error(err))
		return account.ResetPasswordToken{}, err
	}
	var token account.ResetPasswordToken
	err = q.db.Get(&token, sql, args...)
	if err != nil {
		q.log.Info("Fail to get reset password token", zap.Error(err))
		return account.ResetPasswordToken{}, err
	}
	return token, nil
}

func (q *Queries) GetResetPasswordTokenByUserId(ctx context.Context, userId string) (account.ResetPasswordToken, error) {
	sql, args, err := sq.Select("token_id", "user_id", "created_date").From("reset_password_token").Where(sq.Eq{"user_id": userId}).ToSql()
	if err != nil {
		q.log.Warn("Fail to get reset password token query", zap.Error(err))
		return account.ResetPasswordToken{}, err
	}
	var token account.ResetPasswordToken
	err = q.db.Get(&token, sql, args...)
	if err != nil {
		q.log.Info("Fail to get reset password token", zap.Error(err))
		return account.ResetPasswordToken{}, err
	}
	return token, nil
}

func (q *Queries) DeleteResetPasswordToken(ctx context.Context, tokenId string) error {
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
