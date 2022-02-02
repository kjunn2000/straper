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
