import { getLocalStorage } from "../store/localStorage";

var socket;

const connect = cb => {
    var identity = getLocalStorage("identity");
    socket = new WebSocket(`ws://localhost:8080/api/v1/protected/ws/${identity.user_id}`);
    console.log("Connecting");

    socket.onopen = () => {
      console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
      console.log("Successfully receive message")
      const data  = JSON.parse(msg.data)
      cb(data);
    };

    socket.onclose = event => {
      console.log("Socket Closed Connection: ", event);
      socket.close()
    };

    socket.onerror = error => {
      console.log("Socket Error: ", error);
      socket.close()
    };
  }

let sendMsg = (channelId, username, content) => {
    const payload = {
      type : "Message", 
      channel_id : channelId,
      creator_name: username,
      content
    }
    console.log("Sending msg...");
    socket.send(JSON.stringify(payload));
};

export { connect, sendMsg };