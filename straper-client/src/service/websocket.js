import { getLocalStorage } from "../store/localStorage";
import {
  getAsByteArray,
  base64ToArray,
  createAndDownloadBlobFile,
} from "./file";

var socket;

const connect = (pushMsg, pushCard) => {
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
      pushMsg(data.payload);
    } else {
      pushCard(data);
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

let sendMsg = async (type, channelId, creatorId, content) => {
  const payload = {
    type,
    channel_id: channelId,
    creator_id: creatorId,
  };
  if (type === "MESSAGE") {
    payload.content = content;
  } else if (type === "FILE") {
    const result = await getAsByteArray(content);
    payload.file_name = content.name;
    payload.file_type = content.type;
    payload.file_bytes = Array.from(result);
  }
  console.log("Sending msg...");
  const dto = {
    type: "CHAT_MESSAGE",
    payload,
  };
  console.log(dto);

  socket.send(JSON.stringify(dto));
};

export { connect, sendMsg };
