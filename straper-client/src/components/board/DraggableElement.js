import { Draggable, Droppable } from "react-beautiful-dnd";
import ListItem from "./ListItem";
import React, { useState } from "react";
import { AiFillDelete } from "react-icons/ai";
import { iconStyle } from "../../utils/style/icon.js";
import AddComponent from "./AddComponent";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import { DragItem, DragListItem } from "../../utils/style/div";

const DraggableElement = ({ element }) => {
  const board = useBoardStore((state) => state.board);
  const [listName, setListName] = useState(element.list_name);

  const handleListNameUpdate = () => {
    if (element.list_name === listName) {
      return;
    }
    element.list_name = listName;
    sendBoardMsg("BOARD_UPDATE_LIST", board.workspace_id, element);
  };

  return (
    <Draggable draggableId={element.list_id} index={element.order_index}>
      {(provided, snapshot) => {
        return (
          <DragListItem
            ref={provided.innerRef}
            snapshot={snapshot}
            {...provided.draggableProps}
            {...provided.dragHandleProps}
          >
            <div className="rounded-md m-2 bg-gray-600">
              <div className="group flex justify-between text-sm font-medium p-3 bg-gray-700 rounded text-white">
                <span className="font-semibold">
                  <input
                    className="bg-transparent focus:outline-none"
                    defaultValue={listName}
                    onChange={(e) => setListName(e.currentTarget.value)}
                    onBlur={() => handleListNameUpdate()}
                  />
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
          </DragListItem>
        );
      }}
    </Draggable>
  );
};

export default DraggableElement;
