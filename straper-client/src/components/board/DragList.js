import React from "react";
import styled from "styled-components";
import { DragDropContext, Droppable } from "react-beautiful-dnd";
import DraggableElement from "./DraggableElement.js";
import useBoardStore from "../../store/boardStore.js";
import AddComponent from "./AddComponent.js";
import { sendBoardMsg } from "../../service/websocket.js";

const DragDropContextContainer = styled.div`
  padding: 20px;
  border-radius: 6px;
`;

function DragList() {
  const taskLists = useBoardStore((state) => state.taskLists);
  const board = useBoardStore((state) => state.board);
  const orderTaskList = useBoardStore((state) => state.orderTaskList);
  const orderCard = useBoardStore((state) => state.orderCard);

  const onDragEnd = ({ source, destination, type }) => {
    if (!destination) return;

    if (type === "COLUMN") {
      if (source.index !== destination.index) {
        const payload = {
          board_id: board.board_id,
          oldListIndex: source.index,
          newListIndex: destination.index,
        };
        sendBoardMsg("BOARD_ORDER_LIST", board.workspace_id, payload);
        orderTaskList(payload);
      }
      return;
    }

    if (
      source.index !== destination.index ||
      source.droppableId !== destination.droppableId
    ) {
      const payload = {
        sourceListId: source.droppableId,
        destListId: destination.droppableId,
        oldCardIndex: source.index,
        newCardIndex: destination.index,
      };
      sendBoardMsg("BOARD_ORDER_CARD", board.workspace_id, payload);
      orderCard(payload);
    }
  };

  const handleAddNewList = (value) => {
    const payload = {
      list_name: value,
      board_id: board.board_id,
      order_index: taskLists.length + 1,
    };
    sendBoardMsg("BOARD_ADD_LIST", board.workspace_id, payload);
  };

  return !board.board_id ? (
    <svg className="animate-spin h-5 w-5 mr-3 ..." viewBox="0 0 24 24"></svg>
  ) : (
    <DragDropContextContainer className="flex">
      <DragDropContext onDragEnd={onDragEnd}>
        <div className="flex">
          <Droppable
            droppableId={board.board_id}
            direction="horizontal"
            type="COLUMN"
          >
            {(provided) => (
              <div
                {...provided.droppableProps}
                ref={provided.innerRef}
                className="flex"
              >
                {taskLists.map((taskList, i) => (
                  <DraggableElement
                    element={taskList}
                    key={taskList.list_id}
                    index={i}
                  />
                ))}
                {provided.placeholder}
                <AddComponent
                  action={handleAddNewList}
                  type="List"
                  text="+ Add New List"
                />
              </div>
            )}
          </Droppable>
        </div>
      </DragDropContext>
    </DragDropContextContainer>
  );
}

export default DragList;
