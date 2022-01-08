import { getLocalStorage } from "../store/localStorage";
import {
  convertFileToByteArray,
  getAsByteArray,
  base64ToArray,
} from "./object";

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
    console.log(base64ToArray(data.file));
    
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
    const result = await getAsByteArray(content);
    payload.filename = content.name;
    payload.file = Array.from(result);
  }
  console.log("Sending msg...");
  // socket.send(JSON.stringify(payload));
};

export { connect, sendMsg };
