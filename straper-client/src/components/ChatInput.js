import React, { useRef } from "react";
import { sendMsg } from "../service/websocket";
import useIdentifyStore from "../store/identityStore";
import useWorkspaceStore from "../store/workspaceStore";

const ChatInput = ({ scrollToBottom }) => {
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const identity = useIdentifyStore((state) => state.identity);

  const inputRef = useRef(null);

  const defaultPlaceHolder = "Message #" + currChannel?.channel_name;

  const handleKeyDown = (event) => {
    if (event.key !== "Enter") {
      return;
    }
    sendMessage();
  };

  const sendMessage = () => {
    const msg = inputRef.current.value;
    if (!msg || msg === "") {
      return;
    }
    sendMsg(currChannel.channel_id, identity.username, msg);
    inputRef.current.value = "";
    scrollToBottom();
  };

  return (
    <div className="relative flex">
      <div className="p-3 w-full flex">
        <div className="bg-gray-800 bg-opacity-40 rounded-lg w-full">
          <input
            ref={inputRef}
            className="bg-transparent p-3 w-full text-white focus:outline-none"
            placeholder={defaultPlaceHolder}
            onKeyDown={(e) => handleKeyDown(e)}
          />
        </div>
      </div>
      <div className="absolute right-0 items-center inset-y-0 hidden sm:flex">
        <button
          type="button"
          className="inline-flex items-center justify-center rounded-full h-10 w-10"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            className="h-7 w-7 text-gray-600 hover:text-gray-400 transition duration-200"
          >
            <path d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"></path>
          </svg>
        </button>
        <button
          type="button"
          className="inline-flex items-center justify-center rounded-full h-12 w-12 transition duration-500 ease-in-out text-white bg-indigo-500 hover:bg-indigo-400 focus:outline-none"
          onClick={() => sendMessage()}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            className="h-6 w-6 transform rotate-90"
          >
            <path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z"></path>
          </svg>
        </button>
      </div>
    </div>
  );
};

export default ChatInput;
