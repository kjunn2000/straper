package bug

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateIssue(ctx context.Context, issue Issue) (Issue, error)
	AddIssueAttachments(ctx context.Context, issueId string, fileTypes []string, files []*multipart.FileHeader) ([]Attachment, error)
	GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]Issue, error)
	UpdateIssue(ctx context.Context, issue Issue) (Issue, error)
	DeleteIssue(ctx context.Context, issueId string) error
	DeleteIssueAttachment(ctx context.Context, fid string) error
	GetEpicLinkOptions(ctx context.Context, workspaceId string) ([]EpicLinkOption, error)
	GetAssigneeOptions(ctx context.Context, workspaceId string) ([]Assignee, error)
	GetAttachment(ctx context.Context, fid string) ([]byte, error)
}

type service struct {
	log   *zap.Logger
	store Repository
	sc    SeaweedfsClient
}

type SeaweedfsClient interface {
	SaveFile(ctx context.Context, reader io.Reader) (string, error)
	GetFile(ctx context.Context, fid string) ([]byte, error)
	DeleteFile(ctx context.Context, fid string) error
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
	return issue, s.store.CreateIssue(ctx, issue)
}

func (s *service) AddIssueAttachments(ctx context.Context, issueId string, fileTypes []string, files []*multipart.FileHeader) ([]Attachment, error) {
	attachments := []Attachment{}
	for i, fh := range files {
		file, err := fh.Open()
		if err != nil {
			return []Attachment{}, err
		}
		fid, err := s.sc.SaveFile(ctx, file)
		if err != nil {
			return []Attachment{}, err
		}
		a := Attachment{
			Fid:      fid,
			FileName: fh.Filename,
			FileType: fileTypes[i],
			IssueId:  issueId,
		}
		if err := s.store.CreateIssueAttachment(ctx, a); err != nil {
			return []Attachment{}, err
		}
		attachments = append(attachments, a)
	}
	return attachments, nil
}

func (s *service) GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]Issue, error) {
	issues, err := s.store.GetIssuesByWorkspaceId(ctx, workspaceId, limit, offset)
	if err != nil {
		return []Issue{}, err
	}
	for i, issue := range issues {
		attachments, err := s.store.GetIssueAttachments(ctx, issue.IssueId)
		if err != nil {
			return []Issue{}, err
		}
		issue.Attachments = attachments
		issues[i] = issue
	}
	return issues, nil
}

func (s *service) UpdateIssue(ctx context.Context, issue Issue) (Issue, error) {
	if err := s.store.UpdateIssue(ctx, issue); err != nil {
		return Issue{}, err
	}
	return s.getIssueByIssueId(ctx, issue.IssueId)
}

func (s *service) getIssueByIssueId(ctx context.Context, issueId string) (Issue, error) {
	issue, err := s.store.GetIssueByIssueId(ctx, issueId)
	if err != nil {
		return Issue{}, err
	}
	attachments, err := s.store.GetIssueAttachments(ctx, issueId)
	if err != nil {
		return Issue{}, err
	}
	issue.Attachments = attachments
	return issue, nil
}

func (s *service) DeleteIssue(ctx context.Context, issueId string) error {
	attachments, err := s.store.GetIssueAttachments(ctx, issueId)
	if err != nil {
		return err
	}
	for _, a := range attachments {
		if err := s.sc.DeleteFile(ctx, a.Fid); err != nil {
			return err
		}
	}
	return s.store.DeleteIssueAndAttachments(ctx, issueId, attachments)
}

func (s *service) DeleteIssueAttachment(ctx context.Context, fid string) error {
	if err := s.sc.DeleteFile(ctx, fid); err != nil {
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

func (s *service) GetAttachment(ctx context.Context, fid string) ([]byte, error) {
	return s.sc.GetFile(ctx, fid)
}
