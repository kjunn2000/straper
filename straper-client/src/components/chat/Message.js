import React, { useState } from "react";
import { useEffect } from "react/cjs/react.development";
import {
  base64ToArray,
  createBlobFile,
  downloadBlobFile,
} from "../../service/file";
import useIdentityStore from "../../store/identityStore";
import FileMessage from "./FileMessage";
import MessageMenu from "../board/MessageMenu";
import { convertToDateString } from "../../service/object";

const Message = ({ msg, editMsg, deleteMsg }) => {
  const identity = useIdentityStore((state) => state.identity);
  const [file, setFile] = useState({});

  useEffect(() => {
    if (msg.type === "FILE") {
      const blob = createBlobFile(base64ToArray(msg.file_bytes), msg.file_type);
      const src = URL.createObjectURL(blob);
      setFile({
        ...msg,
        blob,
        src,
      });
    }
  }, []);

  const isCreator = () => {
    return identity.user_id === msg.creator_id;
  };

  return (
    <div className="chat-message">
      <div
        className={`flex ${
          isCreator() ? "items-end justify-end" : "items-start justify-start"
        }`}
      >
        <div
          className={`flex flex-col max-w-xs mx-2 group ${
            isCreator() ? "items-end" : "items-start"
          }`}
        >
          {isCreator() ? (
            <MessageMenu
              type={msg.type}
              editMsg={() => msg.type === "MESSAGE" && editMsg(msg.content)}
              deleteMsg={deleteMsg}
            />
          ) : (
            <span className={"inline-block pb-3 text-gray-400 font-semibold"}>
              {msg?.user_detail.username}
            </span>
          )}

          {msg.type === "MESSAGE" ? (
            <p
              className={`px-3 py-2 rounded-lg inline-block text-white max-w-sm break-words ${
                isCreator()
                  ? "rounded-br-none bg-indigo-500"
                  : "rounded-bl-none bg-gray-500 text-white"
              }`}
            >
              {msg.content}
            </p>
          ) : (
            <button onClick={() => downloadBlobFile(file.blob, file.file_name)}>
              {file.file_type && file.file_type.startsWith("image/") ? (
                <img src={file.src} alt={file.file_name} className="rounded" />
              ) : (
                <FileMessage file={file} />
              )}
            </button>
          )}
          <div className="flex flex-col invisible text-gray-400 group-hover:visible transition duration-150">
            <div>{convertToDateString(msg.created_date)}</div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Message;
