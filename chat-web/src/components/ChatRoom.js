import React, { useEffect, useState } from "react";
import ChatInput from "./ChatInput";
import Message from "./Message";
import useWorkspaceStore from "../store/workspaceStore"
import {ReactComponent as Text} from "../asset/img/text.svg"
import useMessageStore from "../store/messageStore";

const ChatRoom = () => {

  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const msgs = useMessageStore((state) => state.messages);

  const emptyMessage = (
    <div className="flex flex-col items-center p-3">
      <span className="pb-5 text-white font-semibold">
        START THE CONVERSATION
      </span>
      <Text/>
    </div>
  )

  const loadMessages = msgs.map((msg) => (
    <Message key={msg.message_id} msg={msg} />
  ));

  return (
    <div
      className="text-white w-full h-full font-medium"
      style={{ background: "rgb(54,57,63)" }}
    >
      <div className="flex flex-col items-stretch h-full">
        <div className="text-xl p-3 mb-6 border border-gray-800">
          # {currChannel.channel_name}
        </div>
        <div className="flex flex-col h-full justify-between">
          {
            msgs.length == 0 ? emptyMessage : <div>{loadMessages}</div>
          }
          <div className="p-3 w-full flex">
            <ChatInput />
          </div>
        </div>
      </div>
    </div>
  );
};

export default ChatRoom;
