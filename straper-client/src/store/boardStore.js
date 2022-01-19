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
  addTaskList: (taskList) => {
    set((state) => {
      const newTaskLists = [...state.taskLists, taskList];
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  deleteTaskList: (listId) => {
    console.log("call you");
    console.log(listId);
    set((state) => {
      console.log(state.taskLists);
      const newTaskLists = state.taskLists
        .filter((taskList) => taskList.list_id !== listId)
        .map((taskList, i) => {
          taskList.order_index = i + 1;
          return taskList;
        });
      console.log(newTaskLists);
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
}));

export default useBoardStore;
