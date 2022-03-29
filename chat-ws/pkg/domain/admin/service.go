package admin

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUser(ctx context.Context, userId string) (User, error)
	GetPaginationUsers(ctx context.Context, param PaginationUsersParam) (PaginationUsersResp, error)
	UpdateUser(ctx context.Context, param UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error

	GetWorkspace(ctx context.Context, workspaceId string) (Workspace, error)
	GetPaginationWorkspaces(ctx context.Context, param PaginationWorkspacesParam) (PaginationWorkspacesResp, error)
}

type PaginationService interface {
	DecodeCursor(encodedCursor string) (res time.Time, uuid string, err error)
	EncodeCursor(t time.Time, uuid string) string
}

type service struct {
	log *zap.Logger
	r   Repository
	ps  PaginationService
}

func NewService(log *zap.Logger, r Repository, ps PaginationService) *service {
	return &service{
		log: log,
		r:   r,
		ps:  ps,
	}
}

func (s *service) GetUser(ctx context.Context, userId string) (User, error) {
	return s.r.GetUser(ctx, userId)
}

func (s *service) GetPaginationUsers(ctx context.Context, param PaginationUsersParam) (PaginationUsersResp, error) {
	users, err := s.r.GetUsersByCursor(ctx, param)
	if err != nil {
		return PaginationUsersResp{}, err
	}
	count, err := s.r.GetUsersCount(ctx, param.SearchStr)
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
	if params.IsPasswdUpdate {
		hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		params.Password = string(hashedBytePassword)
	}
	params.UpdatedDate = time.Now()
	err := s.r.UpdateUserByAdmin(ctx, params)
	if err != nil {
		return s.verigyUserFieldError(err)
	}
	return nil
}

func (us *service) verigyUserFieldError(err error) error {
	fetchField := strings.Split(err.Error(), ".")
	field := fetchField[len(fetchField)-1]
	field = field[:len(field)-1]
	switch field {
	case "username":
		return errors.New("username.registered")
	case "email":
		return errors.New("email.registered")
	case "phone_no":
		return errors.New("phone.no.registered")
	default:
		return err
	}
}

func (s *service) DeleteUser(ctx context.Context, userId string) error {
	return s.r.DeleteUser(ctx, userId)
}

func (s *service) GetWorkspace(ctx context.Context, workspaceId string) (Workspace, error) {
	return s.r.GetWorkspace(ctx, workspaceId)
}

func (s *service) GetPaginationWorkspaces(ctx context.Context, param PaginationWorkspacesParam) (PaginationWorkspacesResp, error) {
	if param.Cursor != "" {
		time, uuid, err := s.ps.DecodeCursor(param.Cursor)
		if err != nil {
			return PaginationWorkspacesResp{}, err
		}
		param.CreatedTime = time
		param.Id = uuid
	}
	workspaces, err := s.r.GetWorkspacesByCursor(ctx, param)
	if err != nil {
		return PaginationWorkspacesResp{}, err
	}
	for i, w := range workspaces {
		w.Cursor = s.ps.EncodeCursor(w.CreatedDate, w.Id)
		workspaces[i] = w
	}
	count, err := s.r.GetWorkspacesCount(ctx, param.SearchStr)
	if err != nil {
		return PaginationWorkspacesResp{}, err
	}
	return PaginationWorkspacesResp{
		Workspaces:      workspaces,
		TotalWorkspaces: count,
	}, nil
}
