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
  updateTaskList: (taskList) => {
    set((state) => {
      const newTaskLists = state.taskLists.map((list) => {
        return list.list_id === taskList.list_id ? taskList : list;
      });
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  deleteTaskList: (listId) => {
    set((state) => {
      const newTaskLists = state.taskLists
        .filter((taskList) => taskList.list_id !== listId)
        .map((taskList, i) => {
          taskList.order_index = i + 1;
          return taskList;
        });
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  orderTaskList: (listIds) => {
    set((state) => {
      const result = state.taskLists.reduce((map, obj) => {
        map[obj.list_id] = obj;
        return map;
      }, {});
      const newTaskLists = listIds
        .map((listId) => result[listId])
        .map((taskList, i) => {
          taskList.order_index = i + 1;
          return taskList;
        });
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  addCard: (card) => {
    set((state) => {
      const newTaskLists = state.taskLists.map((taskList) => {
        if (taskList.list_id === card.list_id) {
          taskList.card_list = taskList.card_list
            ? [...taskList.card_list, card]
            : [card];
        }
        return taskList;
      });
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
}));

export default useBoardStore;
