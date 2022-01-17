import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useBoardStore = create((set) => ({
  board: getLocalStorage("board") || {},
  taskLists: getLocalStorage("taskLists") || [],
  setBoard: (board) => {
    setLocalStorage("board", board);
    set((state) => ({
      board: board,
    }));
  },
  setTaskLists: (taskLists) => {
    setLocalStorage("taskLists", taskLists);
    set((state) => ({
      taskLists: taskLists,
    }));
  },
}));

export default useBoardStore;
