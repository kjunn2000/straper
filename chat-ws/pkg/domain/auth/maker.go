package auth

import "time"

type Maker interface {
	CreateToken(userId string, username string, credentialId string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
