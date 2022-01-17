import React, { useEffect } from "react";
import styled from "styled-components";
import { DragDropContext } from "react-beautiful-dnd";
import DraggableElement from "./DraggableElement.js";
import useBoardStore from "../../store/boardStore.js";
import AddComponent from "./AddComponent.js";

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

const lists = ["todo", "inProgress", "done"];

function DragList() {
  const taskLists = useBoardStore((state) => state.taskLists);

  useEffect(() => {}, []);

  const onDragEnd = (result) => {
    // if (!result.destination) {
    //   return;
    // }
    // const listCopy = { ...elements };
    // const sourceList = listCopy[result.source.droppableId];
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

  return (
    <DragDropContextContainer className="flex">
      <DragDropContext onDragEnd={onDragEnd}>
        <ListGrid>
          {taskLists.map((taskList) => (
            <DraggableElement
              element={taskList.cardList}
              key={taskList.list_id}
            />
          ))}
        </ListGrid>
      </DragDropContext>
      <AddComponent type="List" text="+ Add New List" />
    </DragDropContextContainer>
  );
}

export default DragList;
