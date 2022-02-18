package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/bug"
	"go.uber.org/zap"
)

func (q *Queries) CreateIssue(ctx context.Context, issue bug.Issue) error {
	sql, arg, err := sq.Insert("issue").
		Columns("issue_id", "type", "backlog_priority", "summary", "description", "acceptance_criteria",
			"epic_link", "story_point", "replicate_step", "environment", "workaround", "serverity",
			"label", "assignee", "reporter", "due_time", "status", "created_date").
		Values(issue.IssueId, issue.Type, issue.BacklogPriority, issue.Summary, issue.Description, issue.AcceptanceCriteria,
			issue.EpicLink, issue.StoryPoint, issue.ReplicateStep, issue.Environment, issue.Workaround, issue.Serverity,
			issue.Label, issue.Assignee, issue.Reporter, issue.DueTime, issue.Status, issue.CreatedDate).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create issue query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Info("Failed to create issue to db.", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (q *Queries) GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]bug.Issue, error) {
	var issues []bug.Issue
	sql, arg, err := sq.Select("issue_id", "type", "backlog_priority", "summary", "description", "acceptance_criteria",
		"epic_link", "story_point", "replicate_step", "environment", "workaround", "serverity",
		"label", "assignee", "reporter", "due_time", "status", "created_date").
		From("issue").
		Where(sq.Eq{"workspace_id": workspaceId}).
		OrderBy("created_date desc").Limit(limit).Offset(offset).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []bug.Issue{}, err
	}
	err = q.db.Select(&issues, sql, arg...)
	if err != nil {
		return []bug.Issue{}, err
	}
	return issues, nil
}

func (q *Queries) UpdateIssue(ctx context.Context, issue bug.Issue) error {
	sql, args, err := sq.Update("issue").
		Set("type", issue.Type).
		Set("backlog_priority", issue.BacklogPriority).
		Set("summary", issue.Summary).
		Set("description", issue.Description).
		Set("acceptance_criteria", issue.AcceptanceCriteria).
		Set("epic_link", issue.EpicLink).
		Set("story_point", issue.StoryPoint).
		Set("replicate_step", issue.ReplicateStep).
		Set("environment", issue.Environment).
		Set("workaround", issue.Workaround).
		Set("serverity", issue.Serverity).
		Set("label", issue.Label).
		Set("assignee", issue.Assignee).
		Set("due_time", issue.DueTime).
		Set("status", issue.Status).
		Where(sq.Eq{"issue_id": issue.IssueId}).
		ToSql()
	if err != nil {
		q.log.Info("Failed to create update issue sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update issue.", zap.Error(err))
		return err
	}
	return nil
}

func (s *SQLStore) DeleteIssueAndAttachments(ctx context.Context, issueId string, attachments []bug.Attachment) error {
	err := s.execTx(func(q *Queries) error {
		for _, a := range attachments {
			if err := q.DeleteIssueAttachment(ctx, a.Fid); err != nil {
				return err
			}
		}
		if err := q.DeleteIssue(ctx, issueId); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to delete issue.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteIssue(ctx context.Context, issueId string) error {
	sql, args, err := sq.Delete("issue").
		Where(sq.Eq{"issue_id": issueId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create delete sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete issue.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) CreateIssueAttachment(ctx context.Context, a bug.Attachment) error {
	sql, arg, err := sq.Insert("issue_attachment").
		Columns("fid", "file_name", "file_type", "issue_id").
		Values(a.Fid, a.FileName, a.FileType, a.IssueId).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create issue attachment query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Info("Failed to create issue attachment to db.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetIssueAttachments(ctx context.Context, fid string) ([]bug.Attachment, error) {
	var attacments []bug.Attachment
	sql, arg, err := sq.Select("fid", "file_name", "file_type", "issue_id").
		From("issue_attachment").
		Where(sq.Eq{"fid": fid}).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []bug.Attachment{}, err
	}
	err = q.db.Select(&attacments, sql, arg...)
	if err != nil {
		return []bug.Attachment{}, err
	}
	return attacments, nil
}

func (q *Queries) DeleteIssueAttachment(ctx context.Context, fid string) error {
	sql, args, err := sq.Delete("issue_attachment").
		Where(sq.Eq{"fid": fid}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create delete sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete issue attachment.", zap.Error(err))
		return err
	}
	return nil
}
