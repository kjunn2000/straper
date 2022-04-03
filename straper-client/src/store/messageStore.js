import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useMessageStore = create((set) => ({
  messages: getLocalStorage("messages") || [],
  setMesssages: (messages) => {
    set(() => {
      setLocalStorage("messages", messages);
      return { messages };
    });
  },
  pushMessages: (messages) => {
    set((state) => {
      const newMessages = [...state.messages, ...messages];
      setLocalStorage("messages", newMessages);
      return { messages: newMessages };
    });
  },
  pushMessage: (message) => {
    set((state) => {
      const newMessages = [message, ...state.messages];
      setLocalStorage("messages", newMessages);
      return {
        messages: newMessages,
      };
    });
  },
  clearMessages: () => {
    set(() => {
      setLocalStorage("messages", []);
      return {
        messages: [],
      };
    });
  },
  editMessage: ({ message_id, content }) => {
    set((state) => {
      const newMsgs = state.messages.map((msg) => {
        if (msg.message_id === message_id) {
          msg.content = content;
        }
        return msg;
      });
      setLocalStorage("messages", newMsgs);
      return {
        messages: newMsgs,
      };
    });
  },
  deleteMessage: ({ message_id }) => {
    set((state) => {
      const newMessages = state.messages.filter(
        (msg) => msg.message_id !== message_id
      );
      setLocalStorage("messages", newMessages);
      return {
        messages: newMessages,
      };
    });
  },
}));

export default useMessageStore;
