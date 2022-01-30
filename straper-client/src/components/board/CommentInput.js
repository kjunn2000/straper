import React, { useRef } from "react";
import UploadButton from "../button/UploadButton";
import useIdentityStore from "../../store/identityStore";
import { getAsByteArray } from "../../service/file";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";

const CommentInput = ({ cardId, scrollToTop }) => {
  const inputRef = useRef(null);
  const identity = useIdentityStore((state) => state.identity);
  const board = useBoardStore((state) => state.board);

  const handleKeyDown = (event) => {
    if (event.key !== "Enter") {
      return;
    }
    sendComment("MESSAGE", inputRef.current.value);
  };

  const sendComment = async (type, msg) => {
    if (type === "MESSAGE") {
      msg = msg.trim();
    }
    if (!msg || msg === "") {
      return;
    }
    const payload = {
      type,
      card_id: cardId,
      creator_id: identity.user_id,
    };
    if (type === "MESSAGE") {
      payload.content = msg;
    } else if (type === "FILE") {
      const result = await getAsByteArray(msg);
      payload.file_name = msg.name;
      payload.file_type = msg.type;
      payload.file_bytes = Array.from(result);
    }
    sendBoardMsg("BOARD_CARD_COMMENT", board.workspace_id, payload);
    inputRef.current.value = "";
  };

  return (
    <div className="relative flex flex-col space-y-3 w-4/5">
      <div className="w-full flex">
        <input
          ref={inputRef}
          className="p-3 w-full focus:outline-none bg-gray-200 rounded"
          placeholder="Add a comment..."
          onKeyDown={(e) => handleKeyDown(e)}
        />
      </div>
      <div className="flex items-center">
        <button
          type="button"
          className="items-center rounded transition duration-500 ease-in-out text-white 
          bg-indigo-500 hover:bg-indigo-400 focus:outline-none p-2"
          onClick={() => sendComment("MESSAGE", inputRef.current.value)}
        >
          Save
        </button>
        <UploadButton handleFileAction={(file) => sendComment("FILE", file)} />
      </div>
    </div>
  );
};

export default CommentInput;
