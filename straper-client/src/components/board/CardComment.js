import React, { useEffect, useRef, useState } from "react";
import api from "../../axios/api";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import useCommentStore from "../../store/commentStore";
import Message from "../chat/Message";
import CommentInput from "./CommentInput";

const CardComment = ({ cardId }) => {
  const [isBottom, setIsBottom] = useState(false);
  const [currEditMsgId, setCurrEditMsgId] = useState("");
  const [editedMsg, setEditedMsg] = useState("");

  const comments = useCommentStore((state) => state.comments);
  const commentsRef = useRef(null);
  const commentsStartRef = useRef(null);

  const board = useBoardStore((state) => state.board);
  const clearComments = useCommentStore((state) => state.clearComments);
  const pushComments = useCommentStore((state) => state.pushComments);

  useEffect(() => {
    setIsBottom(false);
    fetchComments(true);
  }, []);

  const fetchComments = (firstTime) => {
    if (isBottom && !firstTime) {
      return;
    } else if (firstTime) {
      clearComments();
    }
    const cursor =
      !firstTime && comments && comments.length > 0
        ? comments[comments.length - 1].cursor
        : "";
    api
      .get(`/protected/board/card/comments/${cardId}?cursor=${cursor}`)
      .then((res) => {
        const fetchedData = res.data.Data;
        if (!fetchedData || fetchedData.length === 0) {
          setIsBottom(true);
          return;
        }
        pushComments(fetchedData);
      });
  };

  const scrollToTop = () => {
    setTimeout(() => {
      if (commentsStartRef.current) {
        commentsStartRef.current.scrollIntoView({
          behavior: "smooth",
          block: "start",
        });
      }
    }, 300);
  };

  const handleScroll = () => {
    if (
      commentsRef.current.scrollHeight -
        Math.ceil(commentsRef.current.scrollTop) ===
      commentsRef.current.clientHeight
    ) {
      fetchComments(false);
    }
  };

  const handleEditComment = (msgId, oriContent) => {
    if (oriContent === editedMsg || editedMsg === "") {
      return;
    }
    const payload = {
      comment_id: msgId,
      content: editedMsg,
    };
    sendBoardMsg("BOARD_CARD_EDIT_COMMENT", board.workspace_id, payload);
    setCurrEditMsgId("");
    setEditedMsg("");
  };

  const handleDeleteComment = (commentId, type, content) => {
    const payload = {
      comment_id: commentId,
      type,
      fid: content,
    };
    sendBoardMsg("BOARD_CARD_DELETE_COMMENT", board.workspace_id, payload);
  };

  return (
    <div className="flex flex-col space-y-5">
      <CommentInput cardId={cardId} scrollToTop={scrollToTop} />
      {comments && comments.length > 0 && (
        <div
          className="h-80 overflow-auto"
          ref={commentsRef}
          onScroll={handleScroll}
        >
          <div ref={commentsStartRef} />
          {comments &&
            comments.map((msg) =>
              msg.comment_id !== currEditMsgId ? (
                <Message
                  key={msg.comment_id}
                  msg={msg}
                  editMsg={() => setCurrEditMsgId(msg.comment_id)}
                  deleteMsg={() =>
                    handleDeleteComment(msg.comment_id, msg.type, msg.content)
                  }
                />
              ) : (
                <div
                  className="flex flex-col items-end justify-end"
                  key={msg.comment_id}
                >
                  <input
                    defaultValue={msg.content}
                    className="p-1 rounded focus:outline-none"
                    onChange={(e) => setEditedMsg(e.target.value)}
                  />
                  <div
                    className="inline-flex rounded-md shadow-sm"
                    role="group"
                  >
                    <button
                      type="button"
                      className="py-2 px-4 text-sm font-medium text-gray-900 rounded-l hover:bg-green-100 hover:text-blue-700 focus:ring-2 focus:ring-green-700 focus:text-w dark:bg-green-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-blue-500 dark:focus:text-white"
                      onClick={() =>
                        handleEditComment(msg.comment_id, msg.content)
                      }
                    >
                      Save
                    </button>
                    <button
                      type="button"
                      className="py-2 px-4 text-sm font-medium text-gray-900 hover:bg-gray-100 hover:text-blue-700 focus:ring-2 focus:ring-green-700 focus:text-blue-700 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-blue-500 dark:focus:text-white"
                      onClick={() => setCurrEditMsgId("")}
                    >
                      Cancel
                    </button>
                  </div>
                </div>
              )
            )}
        </div>
      )}
    </div>
  );
};

export default CardComment;
