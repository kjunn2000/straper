import React, { useRef, useState } from "react";
import UploadButton from "../button/UploadButton";
import { IoSendSharp } from "react-icons/io5";
import { isEmpty } from "../../service/object";

const CommentInput = ({ scrollToBottom }) => {
  const inputRef = useRef(null);

  const handleKeyDown = (event) => {
    if (event.key !== "Enter") {
      return;
    }
  };

  const sendComment = (type, msg) => {
    if (type === "MESSAGE") {
      msg = msg.trim();
    }
    if (!msg || msg === "" || isEmpty(msg)) {
      return;
    }
    // sendMsg(type, currChannel.channel_id, identity.user_id, msg);
    inputRef.current.value = "";
    scrollToBottom();
  };

  return (
    <div className="relative flex w-4/5">
      <div className="p-3 w-full flex">
        <div className="bg-gray-800 bg-opacity-40 rounded-lg w-full">
          <input
            ref={inputRef}
            className="p-3 w-full focus:outline-none bg-gray-200 rounded"
            // placeholder={defaultPlaceHolder}
            onKeyDown={(e) => handleKeyDown(e)}
          />
        </div>
      </div>
      <div className="absolute right-0 items-center inset-y-0 hidden sm:flex">
        <UploadButton handleFileAction={(file) => sendComment("FILE", file)} />
        <button
          type="button"
          className="inline-flex items-center justify-center rounded-full h-12 w-12 transition duration-500 ease-in-out text-white bg-indigo-500 hover:bg-indigo-400 focus:outline-none"
          onClick={() => sendComment("MESSAGE", inputRef.current.value)}
        >
          <IoSendSharp size="25" />
        </button>
      </div>
    </div>
  );
};

export default CommentInput;
