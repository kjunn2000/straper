import React, { useRef } from "react";
import { sendChatMsg } from "../../service/websocket";
import useIdentityStore from "../../store/identityStore";
import useWorkspaceStore from "../../store/workspaceStore";
import UploadButton from "../button/UploadButton";
import { IoSendSharp } from "react-icons/io5";
import { isEmpty } from "../../service/object";
import { getAsByteArray } from "../../service/file";

const ChatInput = () => {
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const identity = useIdentityStore((state) => state.identity);

  const inputRef = useRef(null);

  const defaultPlaceHolder = "Message #" + currChannel?.channel_name;

  const handleKeyDown = (event) => {
    if (event.key !== "Enter") {
      return;
    }
    sendMessage("MESSAGE", inputRef.current.value);
  };

  const sendMessage = async (type, content) => {
    if (type === "MESSAGE") {
      content = content.trim();
    }
    if (!content || content === "" || isEmpty(content)) {
      return;
    }
    const payload = {
      type,
      channel_id: currChannel.channel_id,
      creator_id: identity.user_id,
    };
    if (type === "MESSAGE") {
      payload.content = content;
    } else if (type === "FILE") {
      const result = await getAsByteArray(content);
      payload.file_name = content.name;
      payload.file_type = content.type;
      payload.file_bytes = Array.from(result);
    }
    sendChatMsg("CHAT_ADD_MESSAGE", currChannel.channel_id, payload);
    inputRef.current.value = "";
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
        <UploadButton handleFileAction={(file) => sendMessage("FILE", file)} />
        <button
          type="button"
          className="inline-flex items-center justify-center rounded-full h-12 w-12 transition duration-500 ease-in-out text-white bg-indigo-500 hover:bg-indigo-400 focus:outline-none"
          onClick={() => sendMessage("MESSAGE", inputRef.current.value)}
        >
          <IoSendSharp size="25" />
        </button>
      </div>
    </div>
  );
};

export default ChatInput;
