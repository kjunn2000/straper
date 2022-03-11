package bug

import (
	"time"
)

type Issue struct {
	IssueId            string       `json:"issue_id" db:"issue_id"`
	Type               string       `json:"type" db:"type"`
	BacklogPriority    NullString   `json:"backlog_priority" db:"backlog_priority"`
	Summary            string       `json:"summary" db:"summary"`
	Description        NullString   `json:"description" db:"description"`
	AcceptanceCriteria NullString   `json:"acceptance_criteria" db:"acceptance_criteria"`
	EpicLink           NullString   `json:"epic_link" db:"epic_link"`
	StoryPoint         NullInt64    `json:"story_point" db:"story_point"`
	ReplicateStep      NullString   `json:"replicate_step" db:"replicate_step"`
	Environment        NullString   `json:"environment" db:"environment"`
	Workaround         NullString   `json:"workaround" db:"workaround"`
	Serverity          NullString   `json:"serverity" db:"serverity"`
	Label              NullString   `json:"label" db:"label"`
	Assignee           NullString   `json:"assignee" db:"assignee"`
	Reporter           string       `json:"reporter" db:"reporter"`
	DueTime            NullTime     `json:"due_time" db:"due_time"`
	Status             string       `json:"status" db:"status"`
	WorkspaceId        string       `json:"workspace_id" db:"workspace_id"`
	CreatedDate        time.Time    `json:"created_date" db:"created_date"`
	Attachments        []Attachment `json:"attachments"`
}

type Attachment struct {
	Fid       string `json:"fid" db:"fid"`
	FileName  string `json:"file_name" db:"file_name"`
	FileType  string `json:"file_type" db:"file_type"`
	FileBytes []byte `json:"file_bytes"`
	IssueId   string `json:"issue_id" db:"issue_id"`
}

type EpicLinkOption struct {
	IssueId string `json:"issue_id" db:"issue_id"`
	Summary string `json:"summary" db:"summary"`
}

type Assignee struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
}

type AddIssueAttachmentsParam struct {
	IssueId     string       `json:"issue_id" db:"issue_id"`
	Attachments []Attachment `json:"attachments"`
}
