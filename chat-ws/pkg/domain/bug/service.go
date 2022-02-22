package bug

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateIssue(ctx context.Context, issue Issue) (Issue, error)
	AddIssueAttachments(ctx context.Context, param AddIssueAttachmentsParam) ([]Attachment, error)
	GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]Issue, error)
	UpdateIssue(ctx context.Context, issue Issue) (Issue, error)
	DeleteIssue(ctx context.Context, issueId string) error
	DeleteIssueAttachment(ctx context.Context, fid string) error
	GetEpicLinkOptions(ctx context.Context, workspaceId string) ([]EpicLinkOption, error)
	GetAssigneeOptions(ctx context.Context, workspaceId string) ([]Assignee, error)
}

type service struct {
	log   *zap.Logger
	store Repository
	sc    SeaweedfsClient
}

type SeaweedfsClient interface {
	SaveSeaweedfsFile(ctx context.Context, fileBytes []byte) (string, error)
	GetSeaweedfsFile(ctx context.Context, fid string) ([]byte, error)
	DeleteSeaweedfsFile(ctx context.Context, fid string) error
}

func NewService(log *zap.Logger, store Repository, sc SeaweedfsClient) *service {
	return &service{
		log:   log,
		store: store,
		sc:    sc,
	}
}

func (s *service) CreateIssue(ctx context.Context, issue Issue) (Issue, error) {
	issueId, _ := uuid.NewRandom()
	issue.IssueId = issueId.String()
	issue.CreatedDate = time.Now()
	for i, a := range issue.Attachments {
		fid, err := s.sc.SaveSeaweedfsFile(ctx, a.FileBytes)
		if err != nil {
			return Issue{}, err
		}
		a.Fid = fid
		a.IssueId = issueId.String()
		s.store.CreateIssueAttachment(ctx, a)
		issue.Attachments[i] = a
	}
	return issue, s.store.CreateIssue(ctx, issue)
}

func (s *service) AddIssueAttachments(ctx context.Context, param AddIssueAttachmentsParam) ([]Attachment, error) {
	for i, a := range param.Attachments {
		fid, err := s.sc.SaveSeaweedfsFile(ctx, a.FileBytes)
		if err != nil {
			return []Attachment{}, err
		}
		a.Fid = fid
		a.IssueId = param.IssueId
		if err := s.store.CreateIssueAttachment(ctx, a); err != nil {
			return []Attachment{}, err
		}
		param.Attachments[i] = a
	}
	return param.Attachments, nil
}

func (s *service) GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]Issue, error) {
	issues, err := s.store.GetIssuesByWorkspaceId(ctx, workspaceId, limit, offset)
	if err != nil {
		return []Issue{}, err
	}
	for i, issue := range issues {
		attachments, err := s.getIssueAttachmentsByIssueId(ctx, issue.IssueId)
		if err != nil {
			return []Issue{}, err
		}
		issue.Attachments = attachments
		issues[i] = issue
	}
	return issues, nil
}

func (s *service) getIssueByIssueId(ctx context.Context, issueId string) (Issue, error) {
	issue, err := s.store.GetIssueByIssueId(ctx, issueId)
	if err != nil {
		return Issue{}, err
	}
	attachments, err := s.getIssueAttachmentsByIssueId(ctx, issueId)
	if err != nil {
		return Issue{}, err
	}
	issue.Attachments = attachments
	return issue, nil
}

func (s *service) getIssueAttachmentsByIssueId(ctx context.Context, issueId string) ([]Attachment, error) {
	attachments, err := s.store.GetIssueAttachments(ctx, issueId)
	if err != nil {
		return []Attachment{}, err
	}
	for c, attachment := range attachments {
		bytes, err := s.sc.GetSeaweedfsFile(ctx, attachment.Fid)
		if err != nil {
			return []Attachment{}, err
		}
		attachment.FileBytes = bytes
		attachments[c] = attachment
	}
	return attachments, nil
}

func (s *service) UpdateIssue(ctx context.Context, issue Issue) (Issue, error) {
	if err := s.store.UpdateIssue(ctx, issue); err != nil {
		return Issue{}, err
	}
	return s.getIssueByIssueId(ctx, issue.IssueId)
}

func (s *service) DeleteIssue(ctx context.Context, issueId string) error {
	attachments, err := s.store.GetIssueAttachments(ctx, issueId)
	if err != nil {
		return err
	}
	for _, a := range attachments {
		if err := s.sc.DeleteSeaweedfsFile(ctx, a.Fid); err != nil {
			return err
		}
	}
	return s.store.DeleteIssueAndAttachments(ctx, issueId, attachments)
}

func (s *service) DeleteIssueAttachment(ctx context.Context, fid string) error {
	if err := s.sc.DeleteSeaweedfsFile(ctx, fid); err != nil {
		return err
	}
	if err := s.store.DeleteIssueAttachment(ctx, fid); err != nil {
		return err
	}
	return nil
}

func (s *service) GetEpicLinkOptions(ctx context.Context, workspaceId string) ([]EpicLinkOption, error) {
	return s.store.GetEpicListByWorkspaceId(ctx, workspaceId)
}

func (s *service) GetAssigneeOptions(ctx context.Context, workspaceId string) ([]Assignee, error) {
	return s.store.GetAssigneeListByWorkspaceId(ctx, workspaceId)
}
