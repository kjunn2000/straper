import create from "zustand";
import { addToList, removeFromList } from "../service/board";
import { isEmpty } from "../service/object";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useBoardStore = create((set) => ({
  board: getLocalStorage("board") || {},
  taskListsOrder: getLocalStorage("taskListsOrder") || [],
  taskLists: getLocalStorage("taskLists") || {},
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
      taskList.card_list_order = [];
      taskList.card_list = {};
      const newTaskLists = { ...state.taskLists, [taskList.list_id]: taskList };
      const newTaskListsOrder = [...state.taskListsOrder, taskList.list_id];
      setLocalStorage("taskLists", newTaskLists);
      setLocalStorage("taskListsOrder", newTaskListsOrder);
      return { taskLists: newTaskLists, taskListsOrder: newTaskListsOrder };
    });
  },
  updateTaskList: (taskList) => {
    set((state) => {
      const newTaskList = state.taskLists[taskList.list_id];
      newTaskList.list_name = taskList.list_name;
      const newTaskLists = {
        ...state.taskLists,
        [taskList.list_id]: newTaskList,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  deleteTaskList: (listId) => {
    set((state) => {
      delete state.taskLists[listId];
      const newTaskListsOrder = state.taskListsOrder.filter(
        (id) => id !== listId
      );
      setLocalStorage("taskLists", state.taskLists);
      setLocalStorage("taskListsOrder", newTaskListsOrder);
      return { taskLists: state.taskLists, taskListsOrder: newTaskListsOrder };
    });
  },
  orderTaskList: ({ oldListIndex, newListIndex }) => {
    set((state) => {
      const [removed, result] = removeFromList(
        state.taskListsOrder,
        oldListIndex
      );
      const newTaskListsOrder = addToList(result, newListIndex, removed);
      setLocalStorage("taskListsOrder", newTaskListsOrder);
      return { taskListsOrder: newTaskListsOrder };
    });
  },
  setTaskListsOrder: (taskListsOrder) => {
    set((state) => {
      setLocalStorage("taskListsOrder", taskListsOrder);
      return { taskListsOrder: taskListsOrder };
    });
  },
  addCard: (card) => {
    set((state) => {
      const list = state.taskLists[card.list_id];
      list.card_list[card.card_id] = card;
      list.card_list_order = [...list.card_list_order, card.card_id];
      const newTaskLists = {
        ...state.taskLists,
        [card.list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  orderCard: ({ sourceListId, destListId, oldCardIndex, newCardIndex }) => {
    set((state) => {
      const oldTaskList = state.taskLists[sourceListId];
      let [removed, result] = removeFromList(
        oldTaskList.card_list_order,
        oldCardIndex
      );
      const currObj = oldTaskList.card_list[removed];
      delete oldTaskList.card_list[removed];
      oldTaskList.card_list_order = result;
      if (sourceListId === destListId) {
        oldTaskList.card_list_order = addToList(result, newCardIndex, removed);
        oldTaskList.card_list[removed] = currObj;
        const newTaskLists = {
          ...state.taskLists,
          [sourceListId]: oldTaskList,
        };
        setLocalStorage("taskLists", newTaskLists);
        return { taskLists: newTaskLists };
      }
      const newTaskList = state.taskLists[destListId];
      newTaskList.card_list_order = addToList(
        newTaskList.card_list_order,
        newCardIndex,
        removed
      );
      newTaskList.card_list[removed] = currObj;
      const newTaskLists = {
        ...state.taskLists,
        [sourceListId]: oldTaskList,
        [destListId]: newTaskList,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
}));

export default useBoardStore;
