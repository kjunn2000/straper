package dblog

import (
	"context"

	"go.uber.org/zap"
)

type StatusLogger interface {
	UpdateLastSeen(ctx context.Context, credentialId string) error
}

type statusLogger struct {
	log *zap.Logger
	r   Repository
}

func NewStatusLogger(log *zap.Logger, r Repository) *statusLogger {
	return &statusLogger{
		log: log,
		r:   r,
	}
}

func (sl *statusLogger) UpdateLastSeen(ctx context.Context, credentialId string) error {
	return sl.r.UpdateLastSeen(ctx, credentialId)
}
