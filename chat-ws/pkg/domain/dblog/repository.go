package dblog

import "context"

type Repository interface {
	UpdateLastSeen(ctx context.Context, credentialId string) error
}
