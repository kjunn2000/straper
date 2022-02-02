import useIdentityStore from "../store/identityStore";
import { getLocalStorage } from "../store/localStorage";
import { handleWsBoardMsg } from "./board";
import { handleWsChatMsg } from "./chat";

var socket;

const connect = () => {
  var identity = getLocalStorage("identity");
  socket = new WebSocket(
    `ws://localhost:8080/api/v1/protected/ws/${identity.user_id}`
  );
  console.log("Connecting");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg) => {
    console.log("Successfully receive message");
    const data = JSON.parse(msg.data);
    console.log(data);
    if (data.type.startsWith("CHAT")) {
      handleWsChatMsg(data);
    } else {
      handleWsBoardMsg(data);
    }
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
    socket.close();
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
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
export { connect, sendChatMsg, sendBoardMsg };
