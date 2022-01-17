import { Droppable } from "react-beautiful-dnd";
import ListItem from "./ListItem";
import React from "react";
import styled from "styled-components";
import { AiFillDelete, AiFillEdit } from "react-icons/ai";
import { iconStyle } from "../../utils/style/icon.js";

const ColumnHeader = styled.div`
  text-transform: uppercase;
  margin-bottom: 20px;
`;

const DroppableStyles = styled.div`
  padding: 10px;
  border-radius: 6px;
  background: #d4d4d4;
`;

const DraggableElement = ({ element }) => (
  <DroppableStyles>
    <div className="group flex justify-between text-sm font-medium p-3 text-gray-400 hover:bg-gray-700 rounded hover:text-white">
      <span className="font-semibold">{element.list_name}</span>
      <div className="flex">
        <span className="opacity-0 group-hover:opacity-100 cursor-pointer">
          <AiFillEdit style={iconStyle} />
        </span>
        <span className="opacity-0 group-hover:opacity-100 cursor-pointer pl-3">
          <AiFillDelete style={iconStyle} />
        </span>
      </div>
    </div>
    <Droppable droppableId={element.list_id}>
      {(provided) => (
        <div {...provided.droppableProps} ref={provided.innerRef}>
          {element.card_list.map((item, index) => (
            <ListItem key={item.id} item={item} index={index} />
          ))}
          {provided.placeholder}
        </div>
      )}
    </Droppable>
  </DroppableStyles>
);

export default DraggableElement;
