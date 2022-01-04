import React, { useEffect, useState } from "react";
import ChatInput from "./ChatInput";
import Message from "./Message";
import useWorkspaceStore from "../store/workspaceStore";
import { ReactComponent as Text } from "../asset/img/text.svg";
import useMessageStore from "../store/messageStore";

const ChatRoom = () => {
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const msgs = useMessageStore((state) => state.messages);

  const emptyMessage = (
    <div className="flex flex-col items-center p-3">
      <span className="pb-5 text-white font-semibold">
        START THE CONVERSATION
      </span>
      <Text />
    </div>
  );

  const loadMessages = msgs.map((msg) => (
    <Message key={msg.message_id} msg={msg} />
  ));

  return (
    <div
      className="text-white w-full h-full font-medium overflow-auto"
      style={{ background: "rgb(54,57,63)" }}
    >
      <div className="flex-1 p:2 sm:p-6 justify-between flex flex-col h-screen">
        <div className="flex sm:items-center justify-between py-3 border-b-2 border-gray-500">
          <div className="flex items-center space-x-4">
            <div className="flex flex-col leading-tight">
              <div className="text-xl flex items-center">
                <span className="text-gray-300 mr-3">
                  {currChannel.channel_name}
                </span>
              </div>
            </div>
          </div>
        </div>
        {msgs.length == 0 ? (
          emptyMessage
        ) : (
          <div
            id="messages"
            className="flex flex-col justify-start space-y-4 h-full w-full p-3 overflow-y-auto scrollbar-thumb-blue scrollbar-thumb-rounded scrollbar-track-blue-lighter scrollbar-w-2 scrolling-touch"
          >
            {loadMessages}
          </div>
        )}
        <div className="border-t-2 border-gray-500 px-4 pt-4 mb-2 sm:mb-0">
          <ChatInput />
        </div>
      </div>
    </div>
  );
};

export default ChatRoom;
