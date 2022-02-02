import React, { useEffect, useState } from "react";
import ChatInput from "./ChatInput";
import Message from "./Message";
import useWorkspaceStore from "../../store/workspaceStore";
import { ReactComponent as Text } from "../../asset/img/text.svg";
import useMessageStore from "../../store/messageStore";
import { useRef } from "react/cjs/react.development";
import api from "../../axios/api";
import { sendChatMsg } from "../../service/websocket";

const ChatRoom = () => {
  const [offset, setOffset] = useState(0);
  const [isTop, setIsTop] = useState(false);

  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const msgs = useMessageStore((state) => state.messages);
  const pushMessages = useMessageStore((state) => state.pushMessages);
  const clearMessages = useMessageStore((state) => state.clearMessages);

  const messagesRef = useRef(null);
  const messagesEndRef = useRef(null);

  const [seenMsgs, setSeenMsgs] = useState([]);
  const [unSeenMsgs, setUnSeenMsgs] = useState([]);

  useEffect(() => {
    setOffset(0);
    setIsTop(false);
    fetchMessages(true, 25, 0);
  }, [currChannel]);

  useEffect(() => {
    setTimeout(() => scrollToBottom(), 100);
  }, [msgs]);

  const fetchMessages = (firstTime, limit, offset) => {
    if (isTop && !firstTime) {
      return;
    }
    api
      .get(
        `/protected/channel/${currChannel.channel_id}/messages?limit=${limit}&offset=${offset}`
      )
      .then((res) => {
        const fetchedData = res.data.Data;
        if (fetchedData.length === 0 && !firstTime) {
          setIsTop(true);
          return;
        } else if (firstTime) {
          clearMessages();
          setSeenMsgs([]);
          setUnSeenMsgs([]);
        }
        pushMessages(fetchedData);
        splitSeenMsgs(fetchedData);
        setOffset((offset) => offset + 25);
      });
  };

  const splitSeenMsgs = (msgs) => {
    var date = new Date(currChannel.last_accessed);
    setSeenMsgs(msgs.filter((msg) => new Date(msg.created_date) < date));
    setUnSeenMsgs(msgs.filter((msg) => new Date(msg.created_date) >= date));
  };

  const handleScroll = () => {
    if (messagesRef.current.scrollTop === 0) {
      fetchMessages(false, 25, offset);
    }
  };

  const emptyMessage = (
    <div className="flex flex-col items-center p-3">
      <span className="pb-5 text-white font-semibold">
        START THE CONVERSATION
      </span>
      <Text />
    </div>
  );

  const handleEditMessage = () => {
    console.log("editing message...");
  };

  const handleDeleteMessage = (messageId, type, content) => {
    const payload = {
      message_id: messageId,
      type,
      fid: content,
    };
    sendChatMsg("CHAT_DELETE_MESSAGE", currChannel.channel_id, payload);
  };

  const loadMessages = (msgs) =>
    msgs
      .slice(0)
      .reverse()
      .map((msg) => (
        <Message
          key={msg.message_id}
          msg={msg}
          editMsg={handleEditMessage}
          deleteMsg={() =>
            handleDeleteMessage(msg.message_id, msg.type, msg.content)
          }
        />
      ));

  const scrollToBottom = () => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollIntoView({
        behavior: "smooth",
        block: "end",
      });
    }
  };

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
        {msgs.length === 0 ? (
          emptyMessage
        ) : (
          <div
            id="messages"
            className="flex flex-col justify-start space-y-4 h-full w-full p-3 overflow-y-auto scrollbar-thumb-blue scrollbar-thumb-rounded scrollbar-track-blue-lighter scrollbar-w-2 scrolling-touch"
            ref={messagesRef}
            onScroll={handleScroll}
          >
            {/* {loadMessages(seenMsgs)}
            <div className="text-center text-gray-300">NEW MESSAGES</div>
            {loadMessages(unSeenMsgs)} */}
            {loadMessages(msgs)}
            <div ref={messagesEndRef} />
          </div>
        )}
        <div className="border-t-2 border-gray-500 px-4 pt-4 mb-2 sm:mb-0">
          <ChatInput scrollToBottom={scrollToBottom} />
        </div>
      </div>
    </div>
  );
};

export default ChatRoom;
