package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token.is.invalid")
	ErrExpiredToken = errors.New("token.has.expired")
)

type Payload struct {
	ID           uuid.UUID `json:"id"`
	UserId       string    `json:"user_id"`
	Username     string    `json:"username"`
	CredentialId string    `json:"credential_id"`
	IssuedAt     time.Time `json:"issued_at"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func NewPayload(userId string, username string, credentialId string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:           tokenId,
		UserId:       userId,
		Username:     username,
		CredentialId: credentialId,
		IssuedAt:     time.Now(),
		ExpiredAt:    time.Now().Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if payload.ExpiredAt.Before(time.Now()) {
		return ErrExpiredToken
	}
	return nil
}
