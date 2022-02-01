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

const Message = ({ msg, editMsg, deleteMsg }) => {
  const identity = useIdentityStore((state) => state.identity);
  const [file, setFile] = useState({});

  useEffect(() => {
    if (msg.type == "FILE") {
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

  const zeroPad = (num, places) => String(num).padStart(places, "0");

  const convertToDateString = (timestamp) => {
    var date = new Date(timestamp);
    var dd = date.getDate();
    var mm = date.getMonth() + 1;
    var yy = date.getFullYear();
    var hour = date.getHours();
    var min = date.getMinutes();
    return (
      dd + "/" + mm + "/" + yy + " " + zeroPad(hour, 2) + ":" + zeroPad(min, 2)
    );
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
            <MessageMenu editMsg={editMsg} deleteMsg={deleteMsg} />
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
