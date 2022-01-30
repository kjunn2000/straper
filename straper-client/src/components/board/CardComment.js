import React, { useEffect, useRef } from "react";
import useCommentStore from "../../store/commentStore";
import Message from "../chat/Message";
import CommentInput from "./CommentInput";

const CardComment = ({ cardId }) => {
  const comments = useCommentStore((state) => state.comments);
  const commentsRef = useRef(null);
  const commentsStartRef = useRef(null);

  useEffect(() => {
    setTimeout(() => scrollToTop(), 100);
  }, [comments]);

  const scrollToTop = () => {
    if (commentsStartRef.current) {
      commentsStartRef.current.scrollIntoView({
        behavior: "smooth",
        block: "start",
      });
    }
  };

  const handleScroll = () => {
    // if (messagesRef.current.scrollTop === 0) {
    //   fetchMessages(false, 25, offset);
    // }
  };

  return (
    <div className="flex flex-col space-y-5">
      <CommentInput cardId={cardId} />
      <div
        className="h-80 overflow-auto"
        ref={commentsRef}
        onScroll={handleScroll}
      >
        <div ref={commentsStartRef} />
        {comments &&
          comments.length > 1 &&
          comments.map((msg) => (
            <Message key={msg.comment_id} msg={msg} creatorRight={false} />
          ))}
      </div>
    </div>
  );
};

export default CardComment;
