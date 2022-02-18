import React, { useRef } from "react";
import useIdentityStore from "../../store/identityStore";
import { getAsByteArray } from "../../service/file";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import { IoSendSharp } from "react-icons/io5";
import UploadButton from "../../shared/button/UploadButton";

const CommentInput = ({ cardId }) => {
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
    sendBoardMsg("BOARD_CARD_ADD_COMMENT", board.workspace_id, payload);
    inputRef.current.value = "";
  };

  return (
    <div className="relative flex space-x-2 space-y-2 bg-gray-200 rounded items-center">
      <div className="w-full flex">
        <input
          ref={inputRef}
          className="p-3 w-full focus:outline-none bg-gray-200 rounded"
          placeholder="Add a comment..."
          onKeyDown={(e) => handleKeyDown(e)}
        />
      </div>
      <div className="flex">
        <UploadButton
          handleFileAction={(file) => sendComment("FILE", file)}
          className="text-gray-500"
        />
        <button
          type="button"
          className="inline-flex items-center justify-center rounded-full h-12 w-12 transition duration-500 ease-in-out focus:outline-none"
          onClick={() => sendComment("MESSAGE", inputRef.current.value)}
        >
          <IoSendSharp size="25" className="text-gray-500" />
        </button>
      </div>
    </div>
  );
};

export default CommentInput;
