import React, { useEffect, useState } from "react";
import ChatInput from "./ChatInput";
import Message from "./Message";
import useWorkspaceStore from "../../store/workspaceStore";
import { ReactComponent as Text } from "../../asset/img/text.svg";
import useMessageStore from "../../store/messageStore";
import { useRef } from "react/cjs/react.development";
import api from "../../axios/api";
import { sendChatMsg } from "../../service/websocket";
import { MdDriveFileRenameOutline } from "react-icons/md";
import EditChannelDialog from "../workspace/EditChannelDialog";

const ChatRoom = () => {
  const [isTop, setIsTop] = useState(false);
  const [currEditMsgId, setCurrEditMsgId] = useState("");
  const [editedMsg, setEditedMsg] = useState("");
  const [editChannelDialogOpen, setEditChannelDialogOpen] = useState(false);

  const currChannel = useWorkspaceStore((state) => state.currChannel);

  const msgs = useMessageStore((state) => state.messages);
  const setMessages = useMessageStore((state) => state.setMesssages);
  const pushMessages = useMessageStore((state) => state.pushMessages);
  const clearMessages = useMessageStore((state) => state.clearMessages);

  const messagesRef = useRef(null);
  const messagesEndRef = useRef(null);

  useEffect(() => {
    setIsTop(false);
    fetchMessages(true);
  }, [currChannel]);

  const fetchMessages = async (firstTime) => {
    if (isTop && !firstTime) {
      return;
    } else if (firstTime) {
      clearMessages();
    }
    const cursor =
      !firstTime && msgs && msgs.length > 0 ? msgs[msgs.length - 1].cursor : "";
    const res = await api.get(
      `/protected/channel/${currChannel.channel_id}/messages?cursor=${cursor}`
    );
    const fetchedData = res.data.Data;
    if (!fetchedData || fetchedData.length === 0) {
      setIsTop(true);
      return;
    }
    if (firstTime) {
      setMessages(fetchedData);
      scrollToBottom();
    } else {
      pushMessages(fetchedData);
    }
  };

  const scrollToBottom = () => {
    setTimeout(() => {
      if (messagesEndRef.current) {
        messagesEndRef.current.scrollIntoView({
          behavior: "smooth",
          block: "end",
        });
      }
    }, 300);
  };

  const handleScroll = () => {
    if (messagesRef.current.scrollTop === 0) {
      fetchMessages(false);
    }
  };

  const emptyMessage = (
    <div className="flex flex-col items-center p-3">
      <span className="pb-5 text-white font-semibold">
        START THE CONVERSATION
      </span>
      <Text className="w-auto h-60 lg:h-80 max-w-full max-h-full" />
    </div>
  );

  const handleEditMessage = (msgId, oriContent) => {
    if (oriContent === editedMsg || editedMsg === "") {
      return;
    }
    const payload = {
      message_id: msgId,
      content: editedMsg,
    };
    sendChatMsg("CHAT_EDIT_MESSAGE", currChannel.channel_id, payload);
    setCurrEditMsgId("");
    setEditedMsg("");
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
      .map((msg) =>
        msg.message_id !== currEditMsgId ? (
          <Message
            key={msg.message_id}
            msg={msg}
            editMsg={() => setCurrEditMsgId(msg.message_id)}
            deleteMsg={() =>
              handleDeleteMessage(msg.message_id, msg.type, msg.content)
            }
          />
        ) : (
          <div
            className="flex flex-col items-end justify-end"
            key={msg.message_id}
          >
            <input
              defaultValue={msg.content}
              className="p-1 rounded focus:outline-none text-black"
              onChange={(e) => setEditedMsg(e.target.value)}
            />
            <div className="inline-flex rounded-md shadow-sm" role="group">
              <button
                type="button"
                className="py-2 px-4 text-sm font-medium text-gray-900 rounded-l hover:bg-green-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-green-700 focus:text-w dark:bg-green-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-blue-500 dark:focus:text-white"
                onClick={() => handleEditMessage(msg.message_id, msg.content)}
              >
                Save
              </button>
              <button
                type="button"
                className="py-2 px-4 text-sm font-medium text-gray-900 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-green-700 focus:text-blue-700 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white"
                onClick={() => setCurrEditMsgId("")}
              >
                Cancel
              </button>
            </div>
          </div>
        )
      );

  return (
    <div
      className="text-white w-full h-full font-medium overflow-auto"
      style={{ background: "rgb(54,57,63)" }}
    >
      <div className="flex-1 p-2 sm:p-6 justify-between flex flex-col h-screen">
        <div className="group flex sm:items-center justify-between py-1 border-b-2 border-gray-500">
          <div className="flex items-center space-x-4">
            <div className="flex flex-col leading-tight">
              <div className="text-xl flex items-center">
                <span className="text-gray-300 mr-3 flex items-center">
                  {currChannel.channel_name}
                  <MdDriveFileRenameOutline
                    className="opacity-0 group-hover:opacity-100 hover:cursor-pointer transition duration-150"
                    onClick={() => {
                      setEditChannelDialogOpen(true);
                    }}
                  />
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
            {loadMessages(msgs)}
            <div ref={messagesEndRef} />
          </div>
        )}
        <div className="border-t-2 border-gray-500 px-4 pt-4 mb-2 sm:mb-0">
          <ChatInput scrollToBottom={scrollToBottom} />
        </div>
      </div>
      <EditChannelDialog
        isOpen={editChannelDialogOpen}
        close={() => setEditChannelDialogOpen(false)}
        channel={currChannel}
      />
    </div>
  );
};

export default ChatRoom;
