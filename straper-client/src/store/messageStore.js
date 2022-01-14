import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useMessageStore = create((set) => ({
  messages: getLocalStorage("messages") || [],
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
}));

export default useMessageStore;
