import create from "zustand";
import { addToList, removeFromList } from "../service/board";
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
      currObj.list_id = destListId;
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
  updateCardDetail: (updateCardParams) => {
    set((state) => {
      const list = state.taskLists[updateCardParams.list_id];
      const card = list.card_list[updateCardParams.card_id];
      card.title = updateCardParams.title;
      card.description = updateCardParams.description;
      card.priority = updateCardParams.priority;
      list.card_list[updateCardParams.card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [updateCardParams.list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  updateCardDueDate: ({ list_id, card_id, due_date }) => {
    set((state) => {
      const list = state.taskLists[list_id];
      const card = list.card_list[card_id];
      card.due_date = due_date;
      list.card_list[card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  deleteCard: ({ list_id, card_id }) => {
    set((state) => {
      const taskList = state.taskLists[list_id];
      delete taskList.card_list[card_id];
      const newTaskListOrder = taskList.card_list_order.filter(
        (id) => id !== card_id
      );
      taskList.card_list_order = newTaskListOrder;

      const newTaskLists = {
        ...state.taskLists,
        [list_id]: taskList,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  addMembersToCard: ({ list_id, card_id, member_list }) => {
    set((state) => {
      const list = state.taskLists[list_id];
      const card = list.card_list[card_id];
      if (card.member_list) {
        card.member_list = [...card.member_list, ...member_list];
      } else {
        card.member_list = member_list;
      }
      list.card_list[card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  removeMemberFromCard: ({ list_id, card_id, member_id }) => {
    set((state) => {
      const list = state.taskLists[list_id];
      const card = list.card_list[card_id];
      card.member_list = card.member_list.filter((id) => id !== member_id);
      list.card_list[card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  addChecklistItem: (payload) => {
    set((state) => {
      const list = state.taskLists[payload.list_id];
      const card = list.card_list[payload.card_id];
      if (!card.checklist) {
        card.checklist = [];
      }
      delete payload.list_id;
      card.checklist = [...card.checklist, payload];
      list.card_list[card.card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [list.list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  updateChecklistItem: (payload) => {
    set((state) => {
      const list = state.taskLists[payload.list_id];
      const card = list.card_list[payload.card_id];
      if (!card.checklist) {
        card.checklist = [];
      }
      card.checklist = card.checklist.map((item) => {
        if (item.item_id === payload.item_id) {
          item.content = payload.content;
          item.is_checked = payload.is_checked;
        }
        return item;
      });
      list.card_list[card.card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [list.list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
  deleteChecklistItem: (payload) => {
    set((state) => {
      const list = state.taskLists[payload.list_id];
      const card = list.card_list[payload.card_id];
      card.checklist = card.checklist.filter(
        (item) => item.item_id !== payload.item_id
      );
      list.card_list[card.card_id] = card;
      const newTaskLists = {
        ...state.taskLists,
        [list.list_id]: list,
      };
      setLocalStorage("taskLists", newTaskLists);
      return { taskLists: newTaskLists };
    });
  },
}));

export default useBoardStore;
