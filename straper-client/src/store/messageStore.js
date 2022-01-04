import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useMessageStore = create((set) => ({
  messages: getLocalStorage("messages") || [],
  setMessages: (messages) => {
    setLocalStorage("messages", messages);
    set((state) => ({
      messages: messages,
    }));
  },
  pushMessage: (message) => {
    set((state) => {
      const newMessages = [...state.messages, message];
      setLocalStorage("messages", newMessages);
      return {
        messages: newMessages,
      };
    });
  },
}));

export default useMessageStore;
