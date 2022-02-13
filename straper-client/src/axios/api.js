import axios from "axios";
import useAuthStore from "../store/authStore";

const api = axios.create({
  baseURL: "http://localhost:9090/api/v1",
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use(
  (request) => {
    if (request.url.includes("protected")) {
      const accessToken = useAuthStore.getState().accessToken;
      request.headers["Authorization"] = "Bearer " + accessToken;
    }
    return request;
  },
  (error) => {
    return Promise.reject(error);
  }
);

api.interceptors.response.use(
  (res) => {
    return res;
  },
  async (err) => {
    const originalConfig = err.config;
    if (originalConfig.url !== "/auth/login" && err.response) {
      console.log(err.response);
      console.log(originalConfig);
      if (err.response.status === 403 && !originalConfig._retry) {
        originalConfig._retry = true;

        try {
          const res = await api.post("/auth/refresh-token");

          if (res.data.Success) {
            console.log("success refresh token");
            const accessToken = res.data.Data;
            useAuthStore.getState().setAccessToken(accessToken);
            return api(originalConfig);
          }
          // logOut(true);
        } catch (_error) {
          console.log(_error);
          return Promise.reject(_error);
        }
      }
    }

    return Promise.reject(err);
  }
);

export default api;
