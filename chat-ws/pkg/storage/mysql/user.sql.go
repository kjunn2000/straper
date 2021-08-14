package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"go.uber.org/zap"
)

func (q *Queries) CreateUser(ctx context.Context, params account.CreateUserParam) error {
	sql, arg, err := sq.Insert("user").
		Columns("username", "password", "role", "email", "phone_no", "created_date").
		Values(params.Username, params.Password, params.Role, params.Email, params.PhoneNo, params.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Failed to create insert user query.")
		return err
	}
	res, err := q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to insert record to db.", zap.Error(err))
		return err
	}
	r, err := res.RowsAffected()
	if err != nil {
		q.log.Warn("Failed to read result data.")
		return err
	}
	q.log.Info("Successful create a new user.", zap.Int64("count", r))
	return nil
}

func (q *Queries) GetUserByUserId(ctx context.Context, userId string) (account.User, error) {
	var user account.User
	sta, arg, err := sq.Select("user_id", "username", "password", "role", "email", "phone_no", "created_date").
		From("user").Where(sq.Eq{"user_id": userId}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return account.User{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return account.User{}, err
	}
	q.log.Info("Successful select user ID: ", zap.String("user ID", userId))
	return user, nil
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (auth.User, error) {
	var user auth.User
	sta, arg, err := sq.Select("user_id", "username", "password", "role", "email", "phone_no").
		From("user").Where(sq.Eq{"username": username}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return auth.User{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return auth.User{}, err
	}
	q.log.Info("Successful select username : ", zap.String("username", username))
	return user, nil
}

func (q *Queries) UpdateUser(ctx context.Context, params account.UpdateUserParam) error {
	sql, arg, err := sq.Update("user").
		Set("username", params.Username).
		Set("email", params.Email).
		Set("phone_no", params.PhoneNo).
		Where(sq.Eq{"user_id": params.UserId}).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create update user query.")
		return err
	}
	res, err := q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to update user to db.", zap.Error(err))
		return err
	}
	r, err := res.RowsAffected()
	if err != nil {
		q.log.Warn("Failed to read result data.")
		return err
	}
	q.log.Info("Successful update user.", zap.Int64("count", r))
	return nil
}

func (q *Queries) DeleteUser(ctx context.Context, userId string) error {
	sql, arg, err := sq.Delete("user").Where(sq.Eq{"user_id": userId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create delete user query.")
		return err
	}

	res, err := q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to delete user.", zap.Error(err))
		return err
	}

	r, err := res.RowsAffected()
	if err != nil {
		q.log.Warn("Failed to read result data.")
		return err
	}
	q.log.Info("Successful delete user.", zap.Int64("count", r))
	return nil
}
