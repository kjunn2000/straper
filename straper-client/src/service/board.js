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
    default:
  }
};
