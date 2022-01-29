import React from "react";
import { BsFillChatDotsFill } from "react-icons/bs";
import ChatRoom from "../chat/ChatRoom";
import CommentInput from "./CommentInput";

const CardComment = () => {
  return (
    <div className="flex flex-col">
      <CommentInput />
    </div>
  );
};

export default CardComment;
