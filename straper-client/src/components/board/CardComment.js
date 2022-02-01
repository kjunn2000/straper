import React, { useEffect, useRef, useState } from "react";
import api from "../../axios/api";
import useCommentStore from "../../store/commentStore";
import Message from "../chat/Message";
import CommentInput from "./CommentInput";

const CardComment = ({ cardId }) => {
  const [offset, setOffset] = useState(0);
  const [isBottom, setIsBottom] = useState(false);
  const comments = useCommentStore((state) => state.comments);
  const commentsRef = useRef(null);
  const commentsStartRef = useRef(null);

  const clearComments = useCommentStore((state) => state.clearComments);
  const pushComments = useCommentStore((state) => state.pushComments);

  useEffect(async () => {
    setOffset(0);
    setIsBottom(false);
    await fetchComments(true, 10, 0);
    scrollToTop();
  }, []);

  const scrollToTop = () => {
    if (commentsStartRef.current) {
      commentsStartRef.current.scrollIntoView({
        behavior: "smooth",
        block: "start",
      });
    }
  };

  const fetchComments = async (firstTime, limit, offset) => {
    if (isBottom && !firstTime) {
      return;
    }
    api
      .get(
        `/protected/board/card/comments/${cardId}?limit=${limit}&offset=${offset}`
      )
      .then((res) => {
        const fetchedData = res.data.Data;
        if (!fetchedData && !firstTime) {
          setIsBottom(true);
          return;
        } else if (firstTime) {
          clearComments();
        }
        pushComments(fetchedData);
        setOffset((offset) => offset + 10);
      });
  };

  const handleScroll = () => {
    if (commentsRef.current.scrollTop == commentsRef.current.scrollTopMax) {
      fetchComments(false, 10, offset);
    }
  };

  const handleEditComment = () => {
    console.log("editing comment...");
  };

  const handleDeleteComment = () => {
    console.log("deleting comment...");
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
            <Message
              key={msg.comment_id}
              msg={msg}
              editMsg={handleEditComment}
              deleteMsg={handleDeleteComment}
            />
          ))}
      </div>
    </div>
  );
};

export default CardComment;
