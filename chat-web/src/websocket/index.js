var socket = new WebSocket("ws://localhost:9090/ws");

const connect = cb => {
    console.log("connecting");

    socket.onopen = () => {
      console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
      console.log("Successfully receive message")
      const data  = JSON.parse(msg.data)
      console.log(data)
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

let sendMsg = msg => {
  console.log("sending msg: ", msg);
    const data = {
      type : 3, 
      content: msg,
    }
    socket.send(JSON.stringify(data));
};

export { connect, sendMsg };