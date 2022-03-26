import { Draggable } from "react-beautiful-dnd";
import React from "react";
import Card from "./Card";
import { DragItem } from "../../utils/style/div";

const ListItem = ({ item, index }) => {
  return (
    <Draggable draggableId={item.card_id} index={index}>
      {(provided, snapshot) => {
        return (
          <DragItem
            ref={provided.innerRef}
            snapshot={snapshot}
            {...provided.draggableProps}
            {...provided.dragHandleProps}
            className="w-full"
          >
            <Card card={item} />
          </DragItem>
        );
      }}
    </Draggable>
  );
};

export default ListItem;
