package mysql

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
)

func (q *Queries) CreateUserDetail(ctx context.Context, params CreateUserDetailParam) error {
	sql, arg, err := sq.Insert("user_detail").
		Columns("username", "email", "phone_no", "created_date", "updated_date").
		Values(params.Username, params.Email, params.PhoneNo, params.CreatedDate, params.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Failed to create insert user detail query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Info("Failed to insert record to db.", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (q *Queries) CreateUserCredential(ctx context.Context, params CreateUserCredentialParam) error {
	sql, arg, err := sq.Insert("user_credential").
		Columns("credential_id", "user_id", "password", "role", "status", "created_date", "updated_date").
		Values(params.CredentialId, params.UserId, params.Password, params.Role, params.Status, params.CreatedDate, params.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Failed to create insert user credential query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to insert record to db.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) CreateUserAccessInfo(ctx context.Context, params CreateUserAccessInfo) error {
	sql, arg, err := sq.Insert("user_access_info").
		Columns("credential_id", "last_seen").
		Values(params.CredentialId, params.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Failed to create insert user access info query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to insert record to db.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetUserDetailByUsername(ctx context.Context, username string) (account.UserDetail, error) {
	var user account.UserDetail
	sta, arg, err := sq.Select("user_id", "username", "email", "phone_no", "created_date", "updated_date").
		From("user_detail").Where(sq.Eq{"username": username}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return account.UserDetail{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return account.UserDetail{}, err
	}
	return user, nil
}

func (q *Queries) GetUserDetailByUserId(ctx context.Context, userId string) (account.UserDetail, error) {
	var user account.UserDetail
	sta, arg, err := sq.Select("user_id", "username", "email", "phone_no", "created_date", "updated_date").
		From("user_detail").Where(sq.Eq{"user_id": userId}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return account.UserDetail{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return account.UserDetail{}, err
	}
	return user, nil
}

func (q *Queries) GetUserInfoByUserId(ctx context.Context, userId string) (websocket.UserDetail, error) {
	var user websocket.UserDetail
	sta, arg, err := sq.Select("user_id", "username", "email", "phone_no").
		From("user_detail").Where(sq.Eq{"user_id": userId}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return websocket.UserDetail{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return websocket.UserDetail{}, err
	}
	return user, nil
}

func (q *Queries) GetUserDetailByEmail(ctx context.Context, email string) (account.UserDetail, error) {
	var user account.UserDetail
	sta, arg, err := sq.Select("user_id", "username", "email", "phone_no", "created_date", "updated_date").
		From("user_detail").Where(sq.Eq{"email": email}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return account.UserDetail{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return account.UserDetail{}, err
	}
	return user, nil
}

func (q *Queries) GetUserCredentialByUserId(ctx context.Context, userId string) (auth.User, error) {
	var user auth.User
	sta, arg, err := sq.Select("credential_id", "user_id", "password", "role", "status", "created_date", "updated_date").
		From("user_credential").
		Where(sq.Eq{"user_id": userId}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return auth.User{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return auth.User{}, err
	}
	return user, nil
}

func (q *Queries) GetUserCredentialByUsername(ctx context.Context, username string) (auth.User, error) {
	var user auth.User
	sta, arg, err := sq.Select("credential_id", "user_credential.user_id", "username", "email", "phone_no", "password", "role", "status",
		"user_credential.created_date", "user_credential.updated_date").
		From("user_credential").InnerJoin("user_detail ud on user_credential.user_id = ud.user_id").
		Where(sq.Eq{"username": username}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return auth.User{}, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return auth.User{}, err
	}
	return user, nil
}

func (q *Queries) UpdateUser(ctx context.Context, params account.UpdateUserParam) error {
	sql, arg, err := sq.Update("user_detail").
		Set("username", params.Username).
		Set("email", params.Email).
		Set("phone_no", params.PhoneNo).
		Set("updated_date", params.UpdatedDate).
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

func (q *Queries) UpdateAccountStatus(ctx context.Context, userId, status string) error {
	sql, args, err := sq.Update("user_credential").
		Set("status", status).
		Set("updated_date", time.Now()).
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create user account status query.")
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Warn("Failed to update account status to db.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) UpdateAccountPassword(ctx context.Context, userId, password string) error {
	sql, args, err := sq.Update("user_credential").
		Set("password", password).
		Set("updated_date", time.Now()).
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create update account password query.")
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Warn("Failed to update account password to db.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteUserDetail(ctx context.Context, userId string) error {
	sql, arg, err := sq.Delete("user_detail").Where(sq.Eq{"user_id": userId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create delete user query.")
		return err
	}

	res, err := q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to delete user detail.", zap.Error(err))
		return err
	}

	r, err := res.RowsAffected()
	if err != nil {
		q.log.Warn("Failed to read result data.")
		return err
	}
	q.log.Info("Successful delete user detail.", zap.Int64("count", r))
	return nil
}

func (q *Queries) DeleteUserCredential(ctx context.Context, userId string) error {
	sql, arg, err := sq.Delete("user_credential").Where(sq.Eq{"user_id": userId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create delete user credential query.")
		return err
	}

	res, err := q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to delete user credential.", zap.Error(err))
		return err
	}

	r, err := res.RowsAffected()
	if err != nil {
		q.log.Warn("Failed to read result data.")
		return err
	}
	q.log.Info("Successful delete user credential.", zap.Int64("count", r))
	return nil
}

func (q *Queries) DeleteUserAccessInfo(ctx context.Context, credentialId string) error {
	sql, arg, err := sq.Delete("user_access_info").Where(sq.Eq{"credential_id": credentialId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create delete user access info query.")
		return err
	}

	res, err := q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Warn("Failed to delete user access info.", zap.Error(err))
		return err
	}

	r, err := res.RowsAffected()
	if err != nil {
		q.log.Warn("Failed to read result data.")
		return err
	}
	q.log.Info("Successful delete user access info.", zap.Int64("count", r))
	return nil
}
