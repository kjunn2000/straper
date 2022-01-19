import { Droppable } from "react-beautiful-dnd";
import ListItem from "./ListItem";
import React from "react";
import { AiFillDelete, AiFillEdit } from "react-icons/ai";
import { iconStyle } from "../../utils/style/icon.js";
import AddComponent from "./AddComponent";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";

const DraggableElement = ({ element }) => {
  const board = useBoardStore((state) => state.board);
  return (
    <div>
      <div className="rounded-md m-2 bg-gray-600">
        <div className="group flex justify-between text-sm font-medium p-3 text-gray-400 hover:bg-gray-700 rounded hover:text-white">
          <span className="font-semibold">{element.list_name}</span>
          <div className="flex">
            <span className="opacity-0 group-hover:opacity-100 cursor-pointer">
              <AiFillEdit style={iconStyle} />
            </span>
            <span className="opacity-0 group-hover:opacity-100 cursor-pointer pl-3">
              <AiFillDelete
                style={iconStyle}
                onClick={() =>
                  sendBoardMsg(
                    "BOARD_DELETE_LIST",
                    board.workspace_id,
                    element.list_id
                  )
                }
              />
            </span>
          </div>
        </div>
        <Droppable droppableId={element.list_id}>
          {(provided) => (
            <div {...provided.droppableProps} ref={provided.innerRef}>
              {element.card_list &&
                element.card_list.map((item, index) => (
                  <ListItem key={item.id} item={item} index={index} />
                ))}
              {provided.placeholder}
            </div>
          )}
        </Droppable>
        <AddComponent type="Card" text="Add New Card" />
      </div>
    </div>
  );
};

export default DraggableElement;
