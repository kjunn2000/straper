package board

import "context"

type Service interface {
	GetBoardData(ctx context.Context, workspaceId string)
}
