import { set } from "react-hook-form";
import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useCardStore = create((set) => ({
  cards: getLocalStorage("cards") || {},
  addCard: (cardText, cardId) => {
    set((state) => {
      return {
        cards: {
          ...state.cards,
          [cardId]: {
            text: cardText,
            _id: cardId,
          },
        },
      };
    });
  },

  changeCardText: (cardText, cardId) => {
    set((state) => {
      return {
        cards: {
          ...state.cards,
          [cardId]: {
            text: cardText,
            ...state.cards[cardId],
          },
        },
      };
    });
  },

  deleteCard: (cardId) => {
    set((state) => {
      const { [cardId]: deletedCard, ...restOfCards } = state.cards;
      return {
        cards: restOfCards,
      };
    });
  },
  deleteList: (cardIds) => {
    set((state) => {
      return {
        cards: Object.keys(state)
          .filter((cardId) => !cardIds.includes(cardId))
          .reduce((newState, cardId) => ({
            ...newState,
            [cardId]: state[cardId],
          })),
      };
    });
  },
}));

export default useBoardStore;
