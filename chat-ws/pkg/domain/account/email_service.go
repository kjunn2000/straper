package account

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"text/template"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (us *service) CreateAndSendVerifyEmailToken(ctx context.Context, user UserDetail) error {
	token, err := us.CreateVerifyEmailToken(ctx, user.UserId)
	if err != nil {
		return err
	}
	err = us.sendEmail(ctx, user,
		"template/email/accountVerificationEmail.html",
		"Straper Account Verificataion", token.TokenId)
	if err != nil {
		return err
	}
	return nil
}

func (us *service) CreateVerifyEmailToken(ctx context.Context, userId string) (VerifyEmailToken, error) {
	id, _ := uuid.NewRandom()
	token := VerifyEmailToken{
		TokenId:     id.String(),
		UserId:      userId,
		CreatedDate: time.Now(),
	}
	err := us.ur.CreateVerifyEmailToken(ctx, token)
	if err != nil {
		return VerifyEmailToken{}, err
	}
	return token, nil
}

func (s *service) ValidateVerifyEmailToken(ctx context.Context, tokenId string) error {
	token, err := s.ur.GetVerifyEmailToken(ctx, tokenId)
	if err != nil {
		return err
	}
	err = s.ur.ValidateAccountEmail(ctx, token.UserId, tokenId)
	if err != nil {
		return err
	}
	return nil
}

func (us *service) sendEmail(ctx context.Context, userDetail UserDetail, templatePath, subject, tokenId string) error {

	auth := smtp.PlainAuth("", us.config.SenderEmail, us.config.SenderPassword, us.config.SMTPHost)

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		us.log.Warn("Email template file cannot fetch", zap.Error(err))
		return err
	}

	var body bytes.Buffer

	from := fmt.Sprintf("From: %s", us.config.SenderEmail)
	to := fmt.Sprintf("To: %s", userDetail.Email)
	subject = "Subject: " + subject

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("%s\n%s\n%s\n%s\n\n", from, to, subject, mimeHeaders)))

	t.Execute(&body, struct {
		Username string
		TokenId  string
	}{
		Username: userDetail.Username,
		TokenId:  tokenId,
	})

	err = smtp.SendMail(us.config.SMTPHost+":"+us.config.SMTPPort, auth, us.config.SenderEmail,
		[]string{userDetail.Email}, body.Bytes())
	if err != nil {
		us.log.Warn("Email sending fail", zap.Error(err))
		return err
	}

	return nil
}
