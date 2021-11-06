import React, { useState } from "react";
import ChatInput from "./ChatInput";
import Message from "./Message";

const ChatRoom = ({ channel }) => {
  const [msgs, setMsgs] = useState([
    {
      messasge_id: "M00001",
      username: "Juoann",
      content: "Are you hungry ?",
    },
    {
      messasge_id: "M00002",
      username: "King King",
      content: "Yup, I am hungry now.",
    },
  ]);

  const loadMessages = msgs.map((msg) => (
    <Message key={msg.messasge_id} msg={msg} />
  ));

  return (
    <div
      className="text-white w-full font-medium"
      style={{ background: "rgb(54,57,63)" }}
    >
      {channel === undefined ? (
        <div className="text-center p-5"> CHANNEL NO AVAILABLE</div>
      ) : (
        <div className="flex flex-col items-stretch h-full">
          <div className="text-xl p-3 mb-6 border border-gray-800">
            # {channel.channel_name}
          </div>
          <div className="flex flex-col h-full justify-between">
            <div>{loadMessages}</div>
            <div className="p-3 w-full flex">
              <ChatInput channel={channel} />
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ChatRoom;
