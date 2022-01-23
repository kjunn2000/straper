import useBoardStore from "../store/boardStore";
export const handleWsBoardMsg = (msg) => {
  const boardState = useBoardStore.getState();
  switch (msg.type) {
    case "BOARD_ADD_LIST": {
      boardState.addTaskList(msg.payload);
      break;
    }
    case "BOARD_UPDATE_LIST": {
      boardState.updateTaskList(msg.payload);
      break;
    }
    case "BOARD_DELETE_LIST": {
      boardState.deleteTaskList(msg.payload);
      break;
    }
    case "BOARD_ORDER_LIST": {
      boardState.orderTaskList(msg.payload);
      break;
    }
    case "BOARD_ADD_CARD": {
      boardState.addCard(msg.payload);
      break;
    }
    case "BOARD_UPDATE_CARD_TITLE": {
      boardState.updateCardTitle(msg.payload);
      break;
    }
    case "BOARD_ORDER_CARD": {
      boardState.orderCard(msg.payload);
      break;
    }
    default:
  }
};

export const removeFromList = (list, index) => {
  const result = Array.from(list);
  const [removed] = result.splice(index, 1);
  return [removed, result];
};

export const addToList = (list, index, element) => {
  const result = Array.from(list);
  result.splice(index, 0, element);
  return result;
};
