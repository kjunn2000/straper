import React, { useRef, useState } from "react";
import { BsCardChecklist } from "react-icons/bs";
import { isEmpty } from "../../service/object";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import {
  primaryButtonStyle,
  secondaryButtonStyle,
} from "../../utils/style/button";
import ProgressBar from "./ProgressBar";

const Checklist = ({ show, cardId, listId, checklist }) => {
  const [percentage, setPercentage] = useState(100);
  const [isAddItem, setIsAddItem] = useState(false);
  const board = useBoardStore((state) => state.board);
  const inputRef = useRef(null);

  const handleAddChecklistItem = () => {
    const content = inputRef.current.value;
    if (isEmpty(content)) {
      return;
    }
    const payload = {
      list_id: listId,
      content,
      card_id: cardId,
    };
    sendBoardMsg("BOARD_CARD_ADD_CHECKLIST_ITEM", board.workspace_id, payload);
  };

  const close = () => {
    inputRef.current.target.value = "";
    setIsAddItem(false);
  };

  return (
    <div className={show ? "" : "hidden"}>
      <div className="flex self-center py-3 space-x-3">
        <BsCardChecklist size={30} />
        <span className="font-semibold text-lg">TO DO's</span>
      </div>
      {checklist && <ProgressBar progressPercentage={percentage} />}
      {isAddItem ? (
        <div className="flex flex-col space-y-3">
          <input
            className="rounded bg-gray-200 px-1 py-2 w-3/5"
            ref={inputRef}
          />
          <div>
            <button
              className={primaryButtonStyle}
              onClick={() => handleAddChecklistItem()}
            >
              Add
            </button>
            <button className={secondaryButtonStyle} onClick={() => close()}>
              Cancel
            </button>
          </div>
        </div>
      ) : (
        <div
          className="text-gray-500 hover:cursor-pointer"
          onClick={() => setIsAddItem(true)}
        >
          Add an item
        </div>
      )}
    </div>
  );
};

export default Checklist;
