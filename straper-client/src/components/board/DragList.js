import React, { useEffect } from "react";
import styled from "styled-components";
import { DragDropContext } from "react-beautiful-dnd";
import DraggableElement from "./DraggableElement.js";
import useBoardStore from "../../store/boardStore.js";
import AddComponent from "./AddComponent.js";
import { sendBoardMsg } from "../../service/websocket.js";

const DragDropContextContainer = styled.div`
  padding: 20px;
  border-radius: 6px;
`;

const ListGrid = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-gap: 8px;
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

  const onDragEnd = (result) => {
    if (!result.destination) {
      return;
    }
    console.log(result);
    // const listCopy = { ...elements };
    // const sourceList = listCopy[result.source.droppableId];
    // console.log(sourceList);
    // const [removedElement, newSourceList] = removeFromList(
    //   sourceList,
    //   result.source.index
    // );
    // listCopy[result.source.droppableId] = newSourceList;
    // const destinationList = listCopy[result.destination.droppableId];
    // listCopy[result.destination.droppableId] = addToList(
    //   destinationList,
    //   result.destination.index,
    //   removedElement
    // );
    // setElements(listCopy);
  };

  const handleAddNewList = (value) => {
    const payload = {
      list_name: value,
      board_id: board.board_id,
      order_index: taskLists.length + 1,
    };
    sendBoardMsg("BOARD_ADD_LIST", board.workspace_id, payload);
  };

  return (
    <DragDropContextContainer className="flex">
      <DragDropContext onDragEnd={onDragEnd}>
        <div className="flex">
          {taskLists.map((taskList) => (
            <DraggableElement element={taskList} key={taskList.list_id} />
          ))}
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
