package mysql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"go.uber.org/zap"
)

type CreateUserDetailParam struct {
	Username    string
	Email       string
	PhoneNo     string
	CreatedDate time.Time
}

type CreateUserCredentialParam struct {
	CredentialId string
	UserId       string
	Password     string
	Role         string
	Status       string
	CreatedDate  time.Time
}

type CreateUserAccessInfo struct {
	CredentialId string
	CreatedDate  time.Time
}

func (s *SQLStore) CreateUser(ctx context.Context, params account.CreateUserParam) error {
	err := s.execTx(func(q *Queries) error {
		userDetailParam := CreateUserDetailParam{
			params.Username,
			params.Email,
			params.PhoneNo,
			params.CreatedDate,
		}
		err := q.CreateUserDetail(ctx, userDetailParam)
		if err != nil {
			return err
		}
		userDetail, err := q.GetUserDetailByUsername(ctx, userDetailParam.Username)
		if err != nil {
			return err
		}
		credentialId, _ := uuid.NewRandom()
		userCredentialParam := CreateUserCredentialParam{
			credentialId.String(),
			userDetail.UserId,
			params.Password,
			params.Role,
			params.Status,
			params.CreatedDate,
		}
		err = q.CreateUserCredential(ctx, userCredentialParam)
		if err != nil {
			return err
		}
		err = q.CreateUserAccessInfo(ctx,
			CreateUserAccessInfo{credentialId.String(), params.CreatedDate})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to create a new user.", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (s *SQLStore) DeleteUser(ctx context.Context, userId string) error {
	err := s.execTx(func(q *Queries) error {
		userCredential, err := q.GetUserCredentialByUserId(ctx, userId)
		if err != nil {
			return err
		}
		err = q.DeleteUserAccessInfo(ctx, userCredential.CredentialId)
		if err != nil {
			return err
		}
		err = q.DeleteUserCredential(ctx, userId)
		if err != nil {
			return err
		}
		err = q.DeleteUserDetail(ctx, userId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to create a new user.", zap.Error(err))
		return err
	}
	return nil
}

func (s *SQLStore) ValidateAccountEmail(ctx context.Context, userId, tokenId string) error {
	err := s.execTx(func(q *Queries) error {
		err := q.DeleteVerifyEmailToken(ctx, tokenId)
		if err != nil {
			return err
		}
		err = q.UpdateAccountStatus(ctx, userId, account.StatusActive)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to validate account email.", zap.Error(err))
		return err
	}
	return nil
}

func (s *SQLStore) UpdateUserByAdmin(ctx context.Context, param admin.UpdateUserParam) error {
	err := s.execTx(func(q *Queries) error {
		err := q.UpdateUserDetailByAdmin(ctx, admin.UpdateUserDetailParm{
			UserId:      param.UserId,
			Username:    param.Username,
			Email:       param.Email,
			PhoneNo:     param.PhoneNo,
			UpdatedDate: param.UpdatedDate,
		})
		if err != nil {
			return err
		}
		err = q.UpdateUserCredential(ctx, admin.UpdateCredentialParam{
			UserId:         param.UserId,
			Status:         param.Status,
			IsPasswdUpdate: param.IsPasswdUpdate,
			Password:       param.Password,
			UpdatedDate:    param.UpdatedDate,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to update user.", zap.Error(err))
		return err
	}
	return nil
}
