package admin

import (
	"context"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUser(ctx context.Context, userId string) (User, error)
	GetPaginationUsers(ctx context.Context, limit uint64, cursor string, isNext bool) (PaginationUsersResp, error)
	UpdateUser(ctx context.Context, param UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}

type service struct {
	log *zap.Logger
	r   Repository
}

func NewService(log *zap.Logger, r Repository) *service {
	return &service{
		log: log,
		r:   r,
	}
}

func (s *service) GetUser(ctx context.Context, userId string) (User, error) {
	return s.r.GetUser(ctx, userId)
}

func (s *service) GetPaginationUsers(ctx context.Context, limit uint64, cursor string, isNext bool) (PaginationUsersResp, error) {
	users, err := s.r.GetUsersByCursor(ctx, limit, cursor, isNext)
	if err != nil {
		return PaginationUsersResp{}, err
	}
	count, err := s.r.GetUsersCount(ctx)
	if err != nil {
		return PaginationUsersResp{}, err
	}
	return PaginationUsersResp{
		Users:      users,
		TotalUsers: count,
	}, nil
}

func (s *service) UpdateUser(ctx context.Context, params UpdateUserParam) error {
	params.UpdatedDate = time.Now()
	if params.Password != "" {
		hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		params.Password = string(hashedBytePassword)
	}
	return s.r.UpdateUserByAdmin(ctx, params)
}

func (s *service) DeleteUser(ctx context.Context, userId string) error {
	return s.r.DeleteUser(ctx, userId)
}
