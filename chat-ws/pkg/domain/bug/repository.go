package bug

import "context"

type Repository interface {
	CreateIssue(ctx context.Context, issue Issue) error
	GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]Issue, error)
	UpdateIssue(ctx context.Context, issue Issue) error
	CreateIssueAttachment(ctx context.Context, a Attachment) error
	GetIssueAttachments(ctx context.Context, fid string) ([]Attachment, error)
	DeleteIssueAttachment(ctx context.Context, fid string) error
	DeleteIssueAndAttachments(ctx context.Context, issueId string, attachments []Attachment) error
}
