import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useBoardStore = create((set) => ({
  board: getLocalStorage("board") || {},
  addList: (listId) => {
    set((state) => {
      const newList = [...state.lists, listId];
      setLocalStorage("board", newList);
      return { board: newList };
    });
  },
  moveList: (oldListIndex, newListIndex) => {
    set((state) => {
      const newLists = Array.from(state.lists);
      const [removedList] = newLists.splice(oldListIndex, 1);
      newLists.splice(newListIndex, 0, removedList);
      setLocalStorage("board", newLists);
      return { board: newLists };
    });
  },
  deleteList: (listId) => {
    set((state) => {
      const newLists = state.lists.filter((tmpListId) => tmpListId !== listId);
      setLocalStorage("board", newLists);
      return { board: newLists };
    });
  },
}));

export default useBoardStore;
