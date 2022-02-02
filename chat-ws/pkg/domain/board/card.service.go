package board

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (service *service) handleAddCard(ctx context.Context, bytePayload []byte) ([]byte, error) {
	var card Card
	if err := json.Unmarshal(bytePayload, &card); err != nil {
		return []byte{}, err
	}
	cardId, _ := uuid.NewRandom()
	card.CardId = cardId.String()
	card.Priority = NoPriority
	card.CreatedDate = time.Now()
	card.DueDate = time.Now().Add(time.Hour * 24 * 7)
	if err := service.store.CreateCard(ctx, card); err != nil {
		return []byte{}, err
	}
	newPayload, _ := json.Marshal(card)
	return newPayload, nil
}

func (service *service) handleUpdateCard(ctx context.Context, bytePayload []byte) error {
	var updateCardParams UpdateCardParams
	if err := json.Unmarshal(bytePayload, &updateCardParams); err != nil {
		return err
	}
	return service.store.UpdateCard(ctx, updateCardParams)
}

func (service *service) handleUpdateCardDueDate(ctx context.Context, bytePayload []byte) error {
	var updateCardDueDateParams UpdateCardDueDateParams
	if err := json.Unmarshal(bytePayload, &updateCardDueDateParams); err != nil {
		return err
	}
	return service.store.UpdateCardDueDate(ctx, updateCardDueDateParams)
}

func (service *service) handleDeleteCard(ctx context.Context, bytePayload []byte) error {
	var deleteCardParams DeleteCardParams
	if err := json.Unmarshal(bytePayload, &deleteCardParams); err != nil {
		return err
	}
	return service.store.DeleteCard(ctx, deleteCardParams.CardId)
}

func (service *service) handleOrderCard(ctx context.Context, bytePayload []byte) error {
	var orderCardParams OrderCardParams
	if err := json.Unmarshal(bytePayload, &orderCardParams); err != nil {
		return err
	}
	cardList, err := service.store.GetCardListByListId(ctx, orderCardParams.SourceListId)
	if err != nil {
		return err
	}

	target := cardList[orderCardParams.OldCardIndex]

	cardList = append(cardList[:orderCardParams.OldCardIndex], cardList[orderCardParams.OldCardIndex+1:]...)
	if orderCardParams.SourceListId == orderCardParams.DestListId {
		cardList = append(cardList[:orderCardParams.NewCardIndex],
			append([]Card{target}, cardList[orderCardParams.NewCardIndex:]...)...)
	}
	if err := service.updateCardListOrderIndex(ctx, cardList, orderCardParams.SourceListId); err != nil {
		return err
	}

	if orderCardParams.SourceListId != orderCardParams.DestListId {
		destCardList, err := service.store.GetCardListByListId(ctx, orderCardParams.DestListId)
		if err != nil {
			return err
		}
		destCardList = append(destCardList[:orderCardParams.NewCardIndex],
			append([]Card{target}, destCardList[orderCardParams.NewCardIndex:]...)...)
		if err := service.updateCardListOrderIndex(ctx, destCardList, orderCardParams.DestListId); err != nil {
			return err
		}
	}
	return nil
}

func (service *service) updateCardListOrderIndex(ctx context.Context, cardList []Card, listId string) error {
	for i, card := range cardList {
		if err := service.store.UpdateCardOrder(ctx, card.CardId, i+1, listId, card.ListId != listId); err != nil {
			return err
		}
	}
	return nil
}

func (service *service) handleCardAddMembers(ctx context.Context, bytePayload []byte) error {
	var cardAddMemberParams CardAddMembersParams
	if err := json.Unmarshal(bytePayload, &cardAddMemberParams); err != nil {
		return err
	}
	return service.store.AddUserListToCard(ctx, cardAddMemberParams.CardId, cardAddMemberParams.MemberList)
}

func (service *service) handleCardRemoveMember(ctx context.Context, bytePayload []byte) error {
	var cardRemoveMemberParams CardRemoveMemberParams
	if err := json.Unmarshal(bytePayload, &cardRemoveMemberParams); err != nil {
		return err
	}
	return service.store.DeleteUserFromCard(ctx, cardRemoveMemberParams.CardId, cardRemoveMemberParams.MemberId)
}

func (service *service) handleCardAddChecklistItem(ctx context.Context, bytePayload []byte) ([]byte, error) {
	var cardCheckListItemDto CardChecklistItemDto
	if err := json.Unmarshal(bytePayload, &cardCheckListItemDto); err != nil {
		return []byte{}, err
	}
	itemId, _ := uuid.NewUUID()
	cardCheckListItemDto.ItemId = itemId.String()
	cardCheckListItemDto.IsChecked = false
	err := service.store.CreateChecklistItem(ctx, cardCheckListItemDto)
	if err != nil {
		return []byte{}, err
	}
	newPayload, _ := json.Marshal(cardCheckListItemDto)
	return newPayload, nil
}

func (service *service) handleCardUpdateChecklistItem(ctx context.Context, bytePayload []byte) error {
	var cardCheckListItemDto CardChecklistItemDto
	if err := json.Unmarshal(bytePayload, &cardCheckListItemDto); err != nil {
		return err
	}
	return service.store.UpdateChecklistItem(ctx, cardCheckListItemDto)
}

func (service *service) handleCardDeleteChecklistItem(ctx context.Context, bytePayload []byte) error {
	var cardDeleteChecklistParams CardDeleteChecklistItemParams
	if err := json.Unmarshal(bytePayload, &cardDeleteChecklistParams); err != nil {
		return err
	}
	return service.store.DeleteChecklistItem(ctx, cardDeleteChecklistParams.ItemId)
}

func (s *service) handleBoardAddComment(ctx context.Context, bytePayload []byte) ([]byte, error) {
	var comment CardComment
	if err := json.Unmarshal(bytePayload, &comment); err != nil {
		return []byte{}, err
	}
	newId, _ := uuid.NewRandom()
	comment.CommentId = newId.String()
	comment.CreatedDate = time.Now()
	if comment.Type == TypeFile {
		fid, err := s.sc.SaveSeaweedfsFile(ctx, comment.FileBytes)
		if err != nil {
			return []byte{}, err
		}
		comment.Content = fid
	}
	if err := s.store.CreateCardComment(ctx, &comment); err != nil {
		return []byte{}, err
	}
	userDetail, err := s.store.GetBoardUserInfoByUserId(ctx, comment.CreatorId)
	if err != nil {
		s.log.Warn("Fail to fetch user data.", zap.Error(err))
		return []byte{}, err
	}
	comment.UserDetail = userDetail
	newMsg, err := json.Marshal(comment)
	if err != nil {
		return []byte{}, err
	}
	return newMsg, nil
}

func (service *service) handleCardDeleteComment(ctx context.Context, bytePayload []byte) error {
	var cardDeleteCommentParams CardDeleteCommentParams
	if err := json.Unmarshal(bytePayload, &cardDeleteCommentParams); err != nil {
		return err
	}
	if cardDeleteCommentParams.Type == TypeFile {
		if err := service.sc.DeleteSeaweedfsFile(ctx, cardDeleteCommentParams.Fid); err != nil {
			return err
		}
	}
	return service.store.DeleteCardComment(ctx, cardDeleteCommentParams.CommentId)
}
