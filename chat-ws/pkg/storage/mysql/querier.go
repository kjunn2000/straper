package mysql

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/bug"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
)

type Querier interface {
	UserQueries
	EmailTokenQueries
	ResetPasswordQueries
	WorkspaceQueries
	ChannelQueries
	MessageQueries
	WebsocketQueries
	BoardQueries
	TaskListQueries
	CardQueries
	BugQueries
	LogQueries
}

type UserQueries interface {
	CreateUserDetail(ctx context.Context, params CreateUserDetailParam) error
	CreateUserCredential(ctx context.Context, params CreateUserCredentialParam) error
	CreateUserAccessInfo(ctx context.Context, params CreateUserAccessInfo) error
	GetUserDetailByUsername(ctx context.Context, username string) (account.UserDetail, error)
	GetUserDetailByUserId(ctx context.Context, userId string) (account.UserDetail, error)
	GetUserInfoListByWorkspaceId(ctx context.Context, workspaceId string) ([]account.UserInfo, error)
	GetChatUserInfoByUserId(ctx context.Context, userId string) (chatting.UserDetail, error)
	GetBoardUserInfoByUserId(ctx context.Context, userId string) (board.UserDetail, error)
	GetUserDetailByEmail(ctx context.Context, email string) (account.UserDetail, error)
	GetUserCredentialByUsername(ctx context.Context, username string) (auth.User, error)
	GetUserCredentialByUserId(ctx context.Context, userId string) (auth.User, error)
	UpdateUser(ctx context.Context, params account.UpdateUserParam) error
	UpdateAccountStatus(ctx context.Context, userId, status string) error
	UpdateAccountPassword(ctx context.Context, userId, password string) error
	DeleteUser(ctx context.Context, userId string) error

	// admin
	GetUser(ctx context.Context, userId string) (admin.User, error)
	GetUsersByCursor(ctx context.Context, param admin.PaginationUsersParam) ([]admin.User, error)
	GetUsersCount(ctx context.Context, searchStr string) (int, error)
}

type EmailTokenQueries interface {
	CreateVerifyEmailToken(ctx context.Context, token account.VerifyEmailToken) error
	GetVerifyEmailToken(ctx context.Context, userId string) (account.VerifyEmailToken, error)
	DeleteVerifyEmailToken(ctx context.Context, tokenId string) error
}

type ResetPasswordQueries interface {
	CreateResetPasswordToken(ctx context.Context, params account.ResetPasswordToken) error
	GetResetPasswordToken(ctx context.Context, tokenId string) (account.ResetPasswordToken, error)
	GetResetPasswordTokenByUserId(ctx context.Context, userId string) (account.ResetPasswordToken, error)
	DeleteResetPasswordToken(ctx context.Context, tokenId string) error
}

type WorkspaceQueries interface {
	CreateWorkspace(ctx context.Context, w adding.Workspace) error
	AddUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (listing.Workspace, error)
	GetWorkspacesByUserId(ctx context.Context, userId string) ([]listing.Workspace, error)
	UpdateWorkspace(ctx context.Context, workspace editing.Workspace) error
	DeleteWorkspace(ctx context.Context, id string) error
	RemoveUserFromWorkspace(ctx context.Context, workspaceId, userId string) error

	// admin
	GetWorkspace(ctx context.Context, workspaceId string) (admin.Workspace, error)
	GetWorkspacesByCursor(ctx context.Context, param admin.PaginationWorkspacesParam) ([]admin.WorkspaceSummary, error)
	GetWorkspacesCount(ctx context.Context, searchStr string) (int, error)
}

type ChannelQueries interface {
	CreateChannel(ctx context.Context, channel adding.Channel) error
	AddUserToChannel(ctx context.Context, channelId string, userIdList []string) error
	GetChannelByChannelId(ctx context.Context, channelId string) (listing.Channel, error)
	GetChannelsByUserId(ctx context.Context, userId string) ([]listing.Channel, error)
	GetChannelListByWorkspaceId(ctx context.Context, workspaceId string) ([]listing.Channel, error)
	// GetUserListByChannelId(ctx context.Context, channelId string) ([]chatting.UserData, error)
	GetDefaultChannel(ctx context.Context, workspaceId string) (listing.Channel, error)
	GetDefaultChannelByWorkspaceId(ctx context.Context, workspaceId string) (adding.Channel, error)
	UpdateChannel(ctx context.Context, channel editing.Channel) error
	DeleteChannel(ctx context.Context, channelId string) error
	RemoveUserFromChannel(ctx context.Context, channelId string, userId string) error
}

type MessageQueries interface {
	CreateMessage(ctx context.Context, message *chatting.Message) error
	GetChannelMessages(ctx context.Context, channelId string, param chatting.PaginationMessagesParam) ([]chatting.Message, error)
	GetAllChannelMessages(ctx context.Context, channelId string) ([]chatting.Message, error)
	GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]chatting.Message, error)
	EditMessage(ctx context.Context, params chatting.EditChatMessageParams) error
	DeleteMessage(ctx context.Context, messageId string) error
}

type WebsocketQueries interface {
	GetUserListByChannelId(ctx context.Context, channelId string) ([]websocket.UserData, error)
	GetUserListByWorkspaceId(ctx context.Context, workspaceId string) ([]websocket.UserData, error)
}

type BoardQueries interface {
	CreateBoard(ctx context.Context, board board.TaskBoard) error
	GetTaskBoardByWorkspaceId(ctx context.Context, workspaceId string) (board.TaskBoard, error)
	UpdateTaskBoard(ctx context.Context, board board.TaskBoard) error
}

type TaskListQueries interface {
	CreateTaskList(ctx context.Context, taskList board.TaskList) error
	GetTaskListsByBoardId(ctx context.Context, boardId string) ([]board.TaskList, error)
	UpdateTaskList(ctx context.Context, taskList board.UpdateListParams) error
	UpdateTaskListOrder(ctx context.Context, listId string, orderIndex int) error
	DeleteTaskList(ctx context.Context, listId string) error
}

type CardQueries interface {
	CreateCard(ctx context.Context, card board.Card) error
	GetCardListByListId(ctx context.Context, listId string) ([]board.Card, error)
	UpdateCard(ctx context.Context, params board.UpdateCardParams) error
	UpdateCardDueDate(ctx context.Context, params board.UpdateCardDueDateParams) error
	UpdateCardOrder(ctx context.Context, cardId string, orderIndex int, listId string, updateListId bool) error
	UpdateCardIssueLink(ctx context.Context, cardId string, issueLink board.NullString) error
	DeleteCard(ctx context.Context, cardId string) error

	GetUserFromCard(ctx context.Context, cardId string) ([]string, error)
	AddUserListToCard(ctx context.Context, cardId string, userIdList []string) error
	DeleteUserFromCard(ctx context.Context, cardId, userId string) error

	GetChecklistItemsByCardId(ctx context.Context, cardId string) ([]board.CardChecklistItem, error)
	CreateChecklistItem(ctx context.Context, checklistItem board.CardChecklistItemDto) error
	UpdateChecklistItem(ctx context.Context, checklistItem board.CardChecklistItemDto) error
	DeleteChecklistItem(ctx context.Context, itemId string) error

	CreateCardComment(ctx context.Context, comment *board.CardComment) error
	GetCardComments(ctx context.Context, cardId string, param board.PaginationCommentParam) ([]board.CardComment, error)
	EditCardComment(ctx context.Context, params board.CardEditCommentParams) error
	DeleteCardComment(ctx context.Context, commentId string) error
	GetFileCommentsByCardId(ctx context.Context, cardId string) ([]board.CardComment, error)
	GetFileCommentsByListId(ctx context.Context, listId string) ([]board.CardComment, error)

	GetFidsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
}

type BugQueries interface {
	CreateIssue(ctx context.Context, issue bug.Issue) error
	GetIssuesByWorkspaceId(ctx context.Context, workspaceId string, limit, offset uint64) ([]bug.Issue, error)
	GetIssueByIssueId(ctx context.Context, issueId string) (bug.Issue, error)
	GetAttachmentFidsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
	UpdateIssue(ctx context.Context, issue bug.Issue) error
	DeleteIssue(ctx context.Context, issueId string) error
	CreateIssueAttachment(ctx context.Context, a bug.Attachment) error
	GetIssueAttachments(ctx context.Context, issueId string) ([]bug.Attachment, error)
	DeleteIssueAttachment(ctx context.Context, fid string) error
	GetEpicListByWorkspaceId(ctx context.Context, workspaceId string) ([]bug.EpicLinkOption, error)
	GetAssigneeListByWorkspaceId(ctx context.Context, workspaceId string) ([]bug.Assignee, error)
}

type LogQueries interface {
	UpdateLastSeen(ctx context.Context, credentialId string) error
}
