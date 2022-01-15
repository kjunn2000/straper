import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useBoardStore = create((set) => ({
  board: getLocalStorage("board") || {},
}));

export default useBoardStore;
