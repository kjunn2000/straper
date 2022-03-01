package admin

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Service interface {
	GetPaginationUsers(ctx context.Context, limit uint64, cursor string, isNext bool) ([]User, error)
	UpdateUser(ctx context.Context, param UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}

type service struct {
	log *zap.Logger
	r   Repository
}

func (s *service) GetPaginationUsers(ctx context.Context, limit uint64, cursor string, isNext bool) ([]User, error) {
	return s.r.GetPaginationUsers(ctx, limit, cursor, isNext)
}

func (s *service) UpdateUser(ctx context.Context, params UpdateUserParam) error {
	params.UpdatedDate = time.Now()
	return s.r.UpdateUserByAdmin(ctx, params)
}

func (s *service) DeleteUser(ctx context.Context, userId string) error {
	return s.r.DeleteUser(ctx, userId)
}
