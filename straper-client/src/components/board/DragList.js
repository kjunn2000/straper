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

const removeFromList = (list, index) => {
  const result = Array.from(list);
  const [removed] = result.splice(index, 1);
  return [removed, result];
};

const addToList = (list, index, element) => {
  const result = Array.from(list);
  result.splice(index, 0, element);
  return result;
};

function DragList() {
  const taskLists = useBoardStore((state) => state.taskLists);
  const board = useBoardStore((state) => state.board);
  const orderTaskList = useBoardStore((state) => state.orderTaskList);

  const onDragEnd = (result) => {
    if (
      !result.destination ||
      !result.source ||
      result.destination.index === result.source.index
    ) {
      return;
    }
    const listIds = taskLists.map((taskList) => taskList.list_id);
    const [removedElement, newSourceList] = removeFromList(
      listIds,
      result.source.index - 1
    );
    const newListIds = addToList(
      newSourceList,
      result.destination.index - 1,
      removedElement
    );
    sendBoardMsg("BOARD_ORDER_LIST", board.workspace_id, newListIds);
    orderTaskList(newListIds);
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
                {taskLists.map((taskList) => (
                  <DraggableElement element={taskList} key={taskList.list_id} />
                ))}
                {provided.placeholder}
              </div>
            )}
          </Droppable>
        </div>
      </DragDropContext>
      <AddComponent
        action={handleAddNewList}
        type="List"
        text="+ Add New List"
      />
    </DragDropContextContainer>
  );
}

export default DragList;
