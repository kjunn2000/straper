import { Draggable, Droppable } from "react-beautiful-dnd";
import ListItem from "./ListItem";
import React, { useEffect, useState } from "react";
import { AiFillDelete } from "react-icons/ai";
import { iconStyle } from "../../utils/style/icon.js";
import AddComponent from "./AddComponent";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import { DragListItem } from "../../utils/style/div";
import useIdentityStore from "../../store/identityStore";

const DraggableElement = ({ element, index }) => {
  const identity = useIdentityStore((state) => state.identity);
  const board = useBoardStore((state) => state.board);
  const [listName, setListName] = useState(element.list_name);

  useEffect(() => {
    setListName(element.list_name);
  }, [element.list_name]);

  const handleListNameUpdate = () => {
    if (listName === element.list_name) {
      return;
    }
    element.list_name = listName;
    sendBoardMsg("BOARD_UPDATE_LIST", board.workspace_id, element);
  };

  const handleAddNewCard = (value) => {
    const payload = {
      title: value,
      list_id: element.list_id,
      creator_id: identity.user_id,
      order_index: element.card_list ? element.card_list.length + 1 : 1,
    };
    sendBoardMsg("BOARD_ADD_CARD", board.workspace_id, payload);
  };

  return (
    <Draggable draggableId={element.list_id} index={index}>
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
                    value={listName}
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
                      element.card_list.map((card, i) => (
                        <ListItem key={card.card_id} item={card} index={i} />
                      ))}
                    {provided.placeholder}
                    <AddComponent
                      action={handleAddNewCard}
                      type="Card"
                      text="Add New Card"
                    />
                  </div>
                )}
              </Droppable>
            </div>
          </DragListItem>
        );
      }}
    </Draggable>
  );
};

export default DraggableElement;
