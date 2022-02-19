import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useCommentStore = create((set) => ({
  comments: getLocalStorage("comments") || [],
  pushComments: (comments) => {
    set((state) => {
      const newComments = [...state.comments, ...comments];
      setLocalStorage("comments", newComments);
      return { comments: newComments };
    });
  },
  pushComment: (comment) => {
    set((state) => {
      const newComments = [comment, ...state.comments];
      setLocalStorage("comments", newComments);
      return {
        comments: newComments,
      };
    });
  },
  clearComments: () => {
    set(() => {
      setLocalStorage("comments", []);
      return {
        comments: [],
      };
    });
  },
  updateComment: ({ comment_id, content }) => {
    set((state) => {
      const newComments = state.comments.map((comment) => {
        if (comment.comment_id === comment_id) {
          comment.content = content;
        }
        return comment;
      });
      setLocalStorage("comments", newComments);
      return {
        comments: newComments,
      };
    });
  },
  deleteComment: ({ comment_id }) => {
    set((state) => {
      const newComments = state.comments.filter(
        (comment) => comment.comment_id !== comment_id
      );
      setLocalStorage("comments", newComments);
      return {
        comments: newComments,
      };
    });
  },
}));

export default useCommentStore;
