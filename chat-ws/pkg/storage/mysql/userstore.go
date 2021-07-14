package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"go.uber.org/zap"
)

type userStore struct {
	log *zap.Logger
	db  *sqlx.DB
}

func NewUserStore(log *zap.Logger, db *sqlx.DB) *userStore {
	return &userStore{
		log: log,
		db:  db,
	}
}

func (s *userStore) SaveUser(u *domain.User) error {
	sql, arg, err := sq.Insert("account_credential").
		Columns("username", "password", "role", "created_date").
		Values(u.Username, u.Password, u.Role, u.CreatedDate).ToSql()
	if err != nil {
		s.log.Warn("Failed to create insert statement.")
		return err
	}
	res, err := s.db.Exec(sql, arg...)
	if err != nil {
		s.log.Warn("Failed to insert record to db.", zap.Error(err))
		return err
	}
	r, err := res.RowsAffected()
	if err != nil {
		s.log.Warn("Failed to read result data.")
		return err
	}
	s.log.Info("Successful insert record to db.", zap.Int64("count", r))
	return nil
}

func (s *userStore) FindUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	sta, arg, err := sq.Select("user_id", "username", "password", "role", "created_date").From("account_credential").Where(sq.Eq{"username": username}).Limit(1).ToSql()
	if err != nil {
		s.log.Warn("Unable to create select sql.")
		return nil, err
	}
	err = s.db.Get(&user, sta, arg...)
	if err != nil {
		return nil, err
	}
	s.log.Info("Successful select username : ", zap.String("username", username))
	return &user, nil
}
