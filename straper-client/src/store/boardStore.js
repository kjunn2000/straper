import create from "zustand";
import { addToList, removeFromList } from "../service/board";
import { isEmpty } from "../service/object";
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
  orderTaskList: ({ oldListIndex, newListIndex }) => {
    set((state) => {
      const [removed, result] = removeFromList(state.taskLists, oldListIndex);
      const newTaskLists = addToList(result, newListIndex, removed);
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
  updateCardTitle: (payload) => {
    set((state) => {
      const newTaskLists = state.taskLists.map((taskList) => {
        if (taskList.list_id === payload.list_id) {
          const newCardLists = taskList.card_list.map((card) => {
            if (card.card_id === payload.card_id) {
              card.title = payload.title;
            }
            return card;
          });
          taskList.card_list = newCardLists;
        }
        return taskList;
      });
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  orderCard: ({ sourceListId, destListId, oldCardIndex, newCardIndex }) => {
    set((state) => {
      const isSameList = sourceListId === destListId;
      let removedItem;
      let newTaskLists = state.taskLists.map((taskList) => {
        if (taskList.list_id === sourceListId) {
          let [removed, result] = removeFromList(
            taskList.card_list,
            oldCardIndex
          );
          removedItem = removed;
          if (isSameList) {
            result = addToList(result, newCardIndex, removed);
          }
          result = result.map((card, i) => {
            card.order_index = i + 1;
            return card;
          });
          taskList.card_list = result;
        }
        return taskList;
      });
      if (!isSameList) {
        newTaskLists = state.taskLists.map((taskList) => {
          if (taskList.list_id === destListId) {
            if (isEmpty(taskList.card_list)) {
              taskList.card_list = [removedItem];
            } else {
              taskList.card_list = addToList(
                taskList.card_list,
                newCardIndex,
                removedItem
              );
              taskList.card_list = taskList.card_list.map((card, i) => {
                card.order_index = i + 1;
                return card;
              });
            }
          }
          return taskList;
        });
      }
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
}));

export default useBoardStore;
