import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useAuthStore = create((set) => ({
  accessToken: getLocalStorage("accessToken") || "",
  setAccessToken: (accessToken) => {
    setLocalStorage("accessToken", accessToken);
    set((state) => ({
      accessToken: accessToken,
    }));
  },
  clearAccessToken: () => {
    set((state) => ({
      accessToken: "",
    }));
  },
}));

export default useAuthStore;
