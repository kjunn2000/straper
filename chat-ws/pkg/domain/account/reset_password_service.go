package account

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (us *service) ResetAccountPassword(ctx context.Context, email string) error {
	user, err := us.ur.GetUserDetailByEmail(ctx, email)
	if err != nil {
		return errors.New("email.not.found")
	}
	token, err := us.ur.GetResetPasswordTokenByUserId(ctx, user.UserId)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	} else if err == nil && token.CreatedDate.Add(time.Minute*15).After(time.Now()) {
		return errors.New("password_reset_attempt_in_past_15_min")
	}
	if err := us.ur.DeleteResetPasswordToken(ctx, token.TokenId); err != nil {
		return err
	}
	tokenId, _ := uuid.NewUUID()
	if err := us.ur.CreateResetPasswordToken(ctx, ResetPasswordToken{tokenId.String(), user.UserId, time.Now()}); err != nil {
		return err
	}
	if err := us.sendEmail(ctx, user, "template/email/resetPasswordEmail.html", "Straper Reset Account Password Request", tokenId.String()); err != nil {
		return err
	}
	return nil
}

func (us *service) UpdateAccountPassword(ctx context.Context, params UpdatePasswordParam) error {
	token, err := us.ur.GetResetPasswordToken(ctx, params.TokenId)
	if err == sql.ErrNoRows {
		return errors.New("reset.password.token.not.found")
	} else if err != nil {
		return err
	}
	if token.CreatedDate.Add(time.Minute * 15).Before(time.Now()) {
		return errors.New("reset.password.token.expired")
	} else if err = us.ur.DeleteResetPasswordToken(ctx, token.TokenId); err != nil {
		return err
	}
	hashedPassword, err := BcrptHashPassword(params.Password)
	if err != nil {
		return err
	}

	if err = us.ur.UpdateAccountPassword(ctx, token.UserId, hashedPassword); err != nil {
		return err
	}
	return nil
}
