package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"go.uber.org/zap"
)

func (q *Queries) SaveUser(u *account.User) error {
	sql, arg, err := sq.Insert("user").
		Columns("username", "password", "role", "email", "phone_no", "created_date").
		Values(u.Username, u.Password, u.Role, u.Email, u.PhoneNo, u.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Failed to create insert statement.")
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

func (q *Queries) CheckUsernameExist(username string) (bool, error) {
	var user account.User
	sta, arg, err := sq.Select("*").From("user").Where(sq.Eq{"username": username}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Unable to create select sql.")
		return false, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		q.log.Info("Username doesn not exist.", zap.String("username", username))
		return false, err
	}
	q.log.Info("Username existed.", zap.String("username", username))
	return true, nil
}

func (q *Queries) FindUserByUsername(username string) (*auth.User, error) {
	var user auth.User
	sta, arg, err := sq.Select("user_id", "username", "password", "role").
		From("user").Where(sq.Eq{"username": username}).Limit(1).ToSql()
	if err != nil {
		q.log.Warn("Unable to create select sql.")
		return nil, err
	}
	err = q.db.Get(&user, sta, arg...)
	if err != nil {
		return nil, err
	}
	q.log.Info("Successful select username : ", zap.String("username", username))
	return &user, nil
}
