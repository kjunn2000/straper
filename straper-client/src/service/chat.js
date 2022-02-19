import useMessageStore from "../store/messageStore";

export const handleWsChatMsg = (msg) => {
  const messageState = useMessageStore.getState();

  switch (msg.type) {
    case "CHAT_ADD_MESSAGE": {
      messageState.pushMessage(msg.payload);
      break;
    }
    case "CHAT_EDIT_MESSAGE": {
      messageState.editMessage(msg.payload);
      break;
    }
    case "CHAT_DELETE_MESSAGE": {
      messageState.deleteMessage(msg.payload);
      break;
    }
    default:
  }
};
