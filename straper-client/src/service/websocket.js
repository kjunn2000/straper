import useIdentityStore from "../store/identityStore";
import { getLocalStorage } from "../store/localStorage";
import { handleWsBoardMsg } from "./board";
import { handleWsChatMsg } from "./chat";

var socket;

const isSocketOpen = () => {
  return socket !== undefined && socket.readyState === socket.OPEN;
};

const connect = () => {
  var identity = getLocalStorage("identity");
  socket = new WebSocket(
    `ws://localhost:9090/api/v1/protected/ws/${identity.user_id}`
  );

  socket.onopen = () => {
    // Successfully Connected
  };

  socket.onmessage = (msg) => {
    const data = JSON.parse(msg.data);
    if (data.type.startsWith("CHAT")) {
      handleWsChatMsg(data);
    } else {
      handleWsBoardMsg(data);
    }
  };

  socket.onclose = (event) => {
    // Socket Closed Connection
    socket.close();
  };

  socket.onerror = (error) => {
    // Socket Error
    socket.close();
  };
};

const sendChatMsg = (type, channelId, payload) => {
  const senderId = useIdentityStore.getState().identity.user_id;
  const dto = {
    type,
    channel_id: channelId,
    payload,
    sender_id: senderId,
  };
  console.log(dto);
  socket.send(JSON.stringify(dto));
};

const sendBoardMsg = (type, workspaceId, payload) => {
  const senderId = useIdentityStore.getState().identity.user_id;
  const dto = {
    type,
    workspace_id: workspaceId,
    payload,
    sender_id: senderId,
  };
  console.log(dto);
  socket.send(JSON.stringify(dto));
};

const sendUnregisterMsg = () => {
  if (!socket) {
    return;
  }
  const dto = {
    type: "USER_LEAVE",
  };
  socket.send(JSON.stringify(dto));
};

export { isSocketOpen, connect, sendChatMsg, sendBoardMsg, sendUnregisterMsg };
