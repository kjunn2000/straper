package chatting

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
)

var (
	ChannelGeneral = "General"
	ChannelMessage = "channel-message"
)

type Service interface {
	HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error
	GetBoarcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error)
	GetChannelMessages(ctx context.Context, channelId string, userId string, limit, offset uint64) ([]Message, error)
	DeleteSeaweedfsMessagesByChannelId(ctx context.Context, channelId string) error
	DeleteSeaweedfsMessagesByWorkspaceId(ctx context.Context, workspaceId string) error
}

type PubSub interface {
	SubscribeToChannel(ctx context.Context, channelName string) <-chan *redis.Message
	PublishToChannel(ctx context.Context, channelName string, payload []byte) error
}

type WeedMasterResponse struct {
	Count     int    `json:"count"`
	Fid       string `json:"fid"`
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
}

type WeedVolumeResponse struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	ETag string `json:"eTag"`
}

type WeedVolumeLoopUpResponse struct {
	VolumeOrFileId string     `json:"volumeOrFileId"`
	Locations      []Location `json:"locations"`
}

type WeedUploadFileResponse struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Etag string `json:"eTag"`
}

type Location struct {
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
}

type service struct {
	log   *zap.Logger
	store Repository
}

func NewService(log *zap.Logger, store Repository) *service {
	return &service{
		log:   log,
		store: store,
	}
}

func (s *service) HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error {
	newMsg, err := s.saveAndUpdateMessage(ctx, msg)
	if err != nil {
		s.log.Warn("Fail to save message.", zap.Error(err))
		return err
	}
	if err := publishPubSub(ctx, newMsg); err != nil {
		s.log.Warn("Fail to publish message.", zap.Error(err))
		return err
	}
	return nil
}

func (s *service) saveAndUpdateMessage(ctx context.Context, msg *ws.Message) (*ws.Message, error) {
	newId, err := uuid.NewRandom()
	if err != nil {
		return &ws.Message{}, err
	}
	if msg.MessageType == ws.ChatMessage {
		return s.handleChatMessage(ctx, newId.String(), msg)
	} else if msg.MessageType == ws.BoardCardComment {
		return s.handleBoardComment(ctx, newId.String(), msg)
	}
	return nil, errors.New("invalid.msg.type")
}

func (s *service) handleChatMessage(ctx context.Context, newId string, msg *ws.Message) (*ws.Message, error) {
	message, err := s.convertByteArrayToMessage(ctx, msg)
	if err != nil {
		return &ws.Message{}, err
	}
	message.MessageId = newId
	message.CreatedDate = time.Now()
	if message.Type == File {
		fid, err := s.saveFile(ctx, message.FileBytes)
		if err != nil {
			return &ws.Message{}, err
		}
		message.Content = fid
	}
	if err := s.store.CreateMessage(ctx, &message); err != nil {
		return &ws.Message{}, err
	}
	userDetail, err := s.store.GetUserInfoByUserId(ctx, message.CreatorId)
	if err != nil {
		s.log.Warn("Fail to fetch user data.", zap.Error(err))
		return &ws.Message{}, err
	}
	message.UserDetail = userDetail
	newMsg, err := json.Marshal(message)
	if err != nil {
		return &ws.Message{}, err
	}
	msg.Payload.UnmarshalJSON(newMsg)
	return msg, nil
}

func (s *service) handleBoardComment(ctx context.Context, newId string, msg *ws.Message) (*ws.Message, error) {
	comment, err := s.convertByteArrayToCardComment(ctx, msg)
	if err != nil {
		return &ws.Message{}, err
	}
	comment.CommentId = newId
	comment.CreatedDate = time.Now()
	if comment.Type == File {
		fid, err := s.saveFile(ctx, comment.FileBytes)
		if err != nil {
			return &ws.Message{}, err
		}
		comment.Content = fid
	}
	if err := s.store.CreateCardComment(ctx, &comment); err != nil {
		return &ws.Message{}, err
	}
	newMsg, err := json.Marshal(comment)
	if err != nil {
		return &ws.Message{}, err
	}
	msg.Payload.UnmarshalJSON(newMsg)
	return msg, nil
}

func (s *service) convertByteArrayToMessage(ctx context.Context, msg *ws.Message) (Message, error) {
	bytePayload, err := msg.Payload.MarshalJSON()
	if err != nil {
		return Message{}, err
	}
	var message Message
	if err := json.Unmarshal(bytePayload, &message); err != nil {
		return Message{}, err
	}
	return message, nil
}

func (s *service) convertByteArrayToCardComment(ctx context.Context, msg *ws.Message) (CardComment, error) {
	bytePayload, err := msg.Payload.MarshalJSON()
	if err != nil {
		return CardComment{}, err
	}
	var cardComment CardComment
	if err := json.Unmarshal(bytePayload, &cardComment); err != nil {
		return CardComment{}, err
	}
	return cardComment, nil
}

func (s *service) saveFile(ctx context.Context, fileBytes []byte) (string, error) {
	resp, err := http.Get("http://localhost:9333/dir/assign")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var weedMasterResponse WeedMasterResponse
	json.Unmarshal(body, &weedMasterResponse)
	if err != nil {
		return "", err
	}
	url := "http://" + weedMasterResponse.Url + "/" + weedMasterResponse.Fid
	if err := s.SendMultiPartRequest(fileBytes, url); err != nil {
		return "", err
	}
	return weedMasterResponse.Fid, nil
}

func (s *service) SendMultiPartRequest(fileBytes []byte, url string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("file")
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}
	writer.Close()
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	_, err = client.Do(req)
	if err != nil {
		s.log.Warn("Multipart post request to seaweedfs server failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *service) GetBoarcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error) {
	message, err := s.convertByteArrayToMessage(ctx, msg)
	if err != nil {
		return []ws.UserData{}, err
	}
	return s.store.GetUserListByChannelId(ctx, message.ChannelId)
}

func (s *service) GetChannelMessages(ctx context.Context, channelId string, userId string, limit, offset uint64) ([]Message, error) {
	msgs, err := s.store.GetChannelMessages(ctx, channelId, limit, offset)
	if err == sql.ErrNoRows {
		return []Message{}, errors.New("invalid.channel.id")
	} else if err != nil {
		return []Message{}, err
	}
	for i, msg := range msgs {
		if msg.Type == "FILE" {
			fid := strings.Split(msg.Content, ",")
			resp, err := http.Get("http://localhost:9333/dir/lookup?volumeId=" + fid[0])
			if err != nil {
				s.log.Warn("Seaweedfs look up volume failed", zap.Error(err))
				return []Message{}, err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				s.log.Warn("Read response body failed", zap.Error(err))
				return []Message{}, err
			}
			var weedVolumeLoopUpResponse WeedVolumeLoopUpResponse
			json.Unmarshal(body, &weedVolumeLoopUpResponse)

			resp, err = http.Get("http://" + weedVolumeLoopUpResponse.Locations[0].PublicUrl + "/" + msg.Content)
			if err != nil {
				s.log.Warn("Seaweedfs get file failed", zap.Error(err))
				return []Message{}, err
			}
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				s.log.Warn("Read response body failed", zap.Error(err))
				return []Message{}, err
			}
			msg.FileBytes = body
			msgs[i] = msg
		}
		userDetail, err := s.store.GetUserInfoByUserId(ctx, msg.CreatorId)
		if err != nil {
			return []Message{}, err
		} else {
			msgs[i].UserDetail = userDetail
		}
	}
	if err = s.store.UpdateChannelAccessTime(ctx, channelId, userId); err != nil {
		return []Message{}, err
	}
	return msgs, nil
}

func (s *service) DeleteSeaweedfsMessagesByChannelId(ctx context.Context, channelId string) error {

	msgs, err := s.store.GetAllChannelMessages(ctx, channelId)
	if err != nil {
		return err
	}
	return s.deleteSeaweedfsMessages(ctx, msgs)
}

func (s *service) DeleteSeaweedfsMessagesByWorkspaceId(ctx context.Context, workspaceId string) error {

	msgs, err := s.store.GetAllChannelMessagesByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return err
	}
	return s.deleteSeaweedfsMessages(ctx, msgs)
}

func (s *service) deleteSeaweedfsMessages(ctx context.Context, msgs []Message) error {

	for _, msg := range msgs {
		if msg.Type == "FILE" {
			fid := strings.Split(msg.Content, ",")
			resp, err := http.Get("http://localhost:9333/dir/lookup?volumeId=" + fid[0])
			if err != nil {
				s.log.Warn("Seaweedfs look up volume failed", zap.Error(err))
				return err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				s.log.Warn("Read response body failed", zap.Error(err))
				return err
			}
			var weedVolumeLoopUpResponse WeedVolumeLoopUpResponse
			json.Unmarshal(body, &weedVolumeLoopUpResponse)

			client := &http.Client{}

			req, err := http.NewRequest("DELETE", "http://"+weedVolumeLoopUpResponse.Locations[0].PublicUrl+"/"+msg.Content, nil)
			if err != nil {
				return err
			}

			if _, err = client.Do(req); err != nil {
				return err
			}
		}
	}
	return nil
}
