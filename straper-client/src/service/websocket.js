import { getLocalStorage } from "../store/localStorage";
import {
  getAsByteArray,
  base64ToArray,
  createAndDownloadBlobFile,
} from "./file";

var socket;

const connect = (cb) => {
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
    cb(data);
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

let sendMsg = async (type, channelId, username, content) => {
  const payload = {
    type,
    channel_id: channelId,
    creator_name: username,
  };
  if (type === "MESSAGE") {
    payload.content = content;
  } else if (type === "FILE") {
    console.log(content);
    const result = await getAsByteArray(content);
    const fileMessage = {
      file_name: content.name,
      file_type: content.type,
      bytes: Array.from(result),
    };
    payload.file = fileMessage;
  }
  console.log("Sending msg...");
  console.log(payload);
  socket.send(JSON.stringify(payload));
};

export { connect, sendMsg };
