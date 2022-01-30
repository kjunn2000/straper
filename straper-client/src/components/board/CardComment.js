import React from "react";
import { BsFillChatDotsFill } from "react-icons/bs";
import useCommentStore from "../../store/commentStore";
import ChatRoom from "../chat/ChatRoom";
import Message from "../chat/Message";
import CommentInput from "./CommentInput";

const CardComment = ({ cardId }) => {
  const comments = useCommentStore((state) => state.comments);
  console.log(comments);
  return (
    <div className="flex flex-col">
      <CommentInput cardId={cardId} />
      <div>
        {comments &&
          comments.length > 1 &&
          comments.map((msg) => (
            <Message key={msg.message_id} msg={msg} creatorRight={false} />
          ))}
      </div>
    </div>
  );
};

export default CardComment;
