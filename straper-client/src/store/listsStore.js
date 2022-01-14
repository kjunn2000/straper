import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useListsStore = create((set) => ({
  lists: getLocalStorage("lists") || {},
  addList: (listId, listTitle) => {
    set((state) => {
      const newLists = {
        ...state,
        [listId]: {
          _id: listId,
          title: listTitle,
          cards: [],
        },
      };
      setLocalStorage("lists", newLists);
      return { lists: newLists };
    });
  },

  changeListTitle: (listId, listTitle) => {
    set((state) => {
      const newLists = {
        ...state.lists,
        [listId]: { ...state.lists[listId], title: listTitle },
      };
      return { lists: newLists };
    });
  },

  deleteList: (listId) => {
    set((state) => {
      const { [listId]: deletedList, ...restOfLists } = state.lists;
      return { lists: restOfLists };
    });
  },

  addCard: (listId, cardId) => {
    set((state) => {
      const newLists = {
        ...state.lists,
        [listId]: {
          ...state.lists[listId],
          cards: [...state.lists[listId].cards, cardId],
        },
      };
      return { lists: newLists };
    });
  },

  moveCard: (oldCardIndex, newCardIndex, sourceListId, destListId) => {
    set((state) => {
      if (sourceListId === destListId) {
        const newCards = Array.from(state.lists[sourceListId].cards);
        const [removedCard] = newCards.splice(oldCardIndex, 1);
        newCards.splice(newCardIndex, 0, removedCard);
        return {
          lists: {
            ...state.lists,
            [sourceListId]: { ...state.lists[sourceListId], cards: newCards },
          },
        };
      }
      const sourceCards = Array.from(state.lists[sourceListId].cards);
      const [removedCard] = sourceCards.splice(oldCardIndex, 1);
      const destinationCards = Array.from(state.lists[destListId].cards);
      destinationCards.splice(newCardIndex, 0, removedCard);
      return {
        lists: {
          ...state.lists,
          [sourceListId]: { ...state.lists[sourceListId], cards: sourceCards },
          [destListId]: { ...state.lists[destListId], cards: destinationCards },
        },
      };
    });
  },
  deleteCard: (deletedCardId, listId) => {
    set((state) => {
      const filterDeleted = (cardId) => cardId !== deletedCardId;
      return {
        lists: {
          ...state.lists,
          [listId]: {
            ...state.lists[listId],
            cards: state.lists[listId].cards.filter(filterDeleted),
          },
        },
      };
    });
  },
}));
