import React, { useRef, useState, useEffect } from "react";
import { BsCardChecklist } from "react-icons/bs";
import { isEmpty } from "../../service/object";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import {
  primaryButtonStyle,
  secondaryButtonStyle,
} from "../../utils/style/button";
import ProgressBar from "./ProgressBar";
import { IoIosRemoveCircleOutline } from "react-icons/io";

const Checklist = ({ show, cardId, listId, checklist }) => {
  const [percentage, setPercentage] = useState(100);
  const [isAddItem, setIsAddItem] = useState(false);
  const board = useBoardStore((state) => state.board);
  const inputRef = useRef(null);

  useEffect(() => {
    if (!checklist) {
      return;
    }
    const checkedItemCount = checklist.reduce(
      (count, item) => (item.is_checked ? count + 1 : count),
      0
    );
    const pct = (checkedItemCount / checklist.length) * 100;
    setPercentage(pct.toFixed(2));
  }, [checklist]);

  const close = () => {
    inputRef.current.value = "";
    setIsAddItem(false);
  };
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
    close();
  };

  const handleUpdateChecklistItem = (itemId, content, checked) => {
    const payload = {
      list_id: listId,
      item_id: itemId,
      content,
      is_checked: checked,
      card_id: cardId,
    };
    sendBoardMsg(
      "BOARD_CARD_UPDATE_CHECKLIST_ITEM",
      board.workspace_id,
      payload
    );
  };

  const handleDeleteChecklistItem = (itemId) => {
    const payload = {
      list_id: listId,
      card_id: cardId,
      item_id: itemId,
    };
    sendBoardMsg(
      "BOARD_CARD_DELETE_CHECKLIST_ITEM",
      board.workspace_id,
      payload
    );
  };

  return (
    <div className={show ? "" : "hidden"}>
      <div className="flex self-center py-3 space-x-3">
        <BsCardChecklist size={30} />
        <span className="font-semibold text-lg">TO DO's</span>
      </div>
      {checklist && checklist.length > 0 && (
        <div className="flex flex-col space-y-3 p-2">
          <ProgressBar progressPercentage={percentage} />
          {checklist.map((item) => (
            <div
              key={item.item_id}
              className="group w-2/5 grid grid-cols-6 place-items-center"
            >
              <input
                type="checkbox"
                className="col-span-1 rounded hover:cursor-pointer"
                onChange={(e) =>
                  handleUpdateChecklistItem(
                    item.item_id,
                    item.content,
                    e.target.checked
                  )
                }
                checked={item.is_checked}
              />
              <input
                defaultValue={item.content}
                className={
                  "col-span-4 rounded p-1 bg-gray-200 " +
                  (item.is_checked ? "line-through" : "")
                }
                onBlur={(e) => {
                  if (e.target.value === item.content) {
                    return;
                  }
                  handleUpdateChecklistItem(
                    item.item_id,
                    e.target.value,
                    item.is_checked
                  );
                }}
              />
              <span className="col-span-1">
                <IoIosRemoveCircleOutline
                  size={30}
                  className="opacity-0 text-red-600 group-hover:opacity-100 hover:cursor-pointer"
                  onClick={() => handleDeleteChecklistItem(item.item_id)}
                />
              </span>
            </div>
          ))}
        </div>
      )}
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
