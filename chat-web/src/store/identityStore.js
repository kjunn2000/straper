import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useIdentifyStore = create((set) => ({
  identity: getLocalStorage("identity") || {},
  setIdentity: (identity) => {
    set((state) => {
      setLocalStorage("identity", identity);
      return { identity };
    });
  },
  clearIdentity: () => {
    set((state) => ({
      identity: {},
    }));
  },
}));

export default useIdentifyStore;
