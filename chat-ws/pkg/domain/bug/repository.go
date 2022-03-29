package bug

import (
	"context"
)

type Repository interface {
	CreateIssue(ctx context.Context, issue Issue) error
	GetIssuesByWorkspaceId(ctx context.Context, workspaceId string) ([]Issue, error)
	GetIssueByIssueId(ctx context.Context, issueId string) (Issue, error)
	UpdateIssue(ctx context.Context, issue Issue) error
	CreateIssueAttachment(ctx context.Context, a Attachment) error
	GetIssueAttachments(ctx context.Context, issueId string) ([]Attachment, error)
	DeleteIssueAttachment(ctx context.Context, fid string) error
	DeleteIssueAndAttachments(ctx context.Context, issueId string, attachments []Attachment) error
	GetEpicListByWorkspaceId(ctx context.Context, workspaceId string) ([]EpicLinkOption, error)
	GetAssigneeListByWorkspaceId(ctx context.Context, workspaceId string) ([]Assignee, error)
}
