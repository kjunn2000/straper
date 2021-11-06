import create from "zustand";

const useAuthStore = create((set) => ({
  accessToken: window.localStorage.getItem("accessToken") || "",
  setAccessToken: (accessToken) => {
    window.localStorage.setItem("accessToken", accessToken);
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
