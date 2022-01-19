import { Draggable } from "react-beautiful-dnd";
import React from "react";
import Card from "./Card";
import { DragItem } from "../../utils/style/div";

const ListItem = ({ item, index }) => {
  return (
    <Draggable draggableId={item.id} index={index}>
      {(provided, snapshot) => {
        return (
          <DragItem
            ref={provided.innerRef}
            snapshot={snapshot}
            {...provided.draggableProps}
            {...provided.dragHandleProps}
          >
            <Card />
          </DragItem>
        );
      }}
    </Draggable>
  );
};

export default ListItem;
