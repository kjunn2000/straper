import axios from "axios";
import useAuthStore from "../store/authStore";

const instance = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  headers: {
    "Content-Type": "application/json",
  },
});

instance.interceptors.request.use(
	request => {
		if (request.url.includes('protected')){
			const accessToken = useAuthStore.getState().accessToken
			request.headers['Authorization'] = accessToken;
		}
		return request;
	},
	error => {
		return Promise.reject(error);
	}
)

instance.interceptors.response.use(
  (res) => {
    return res;
  },
  async (err) => {
    const originalConfig = err.config;

    if (originalConfig.url !== "/auth/login" && err.response) {

      if (err.response.status === 401 && !originalConfig._retry) {
        originalConfig._retry = true;

        try {
          const rs = await instance.post("/auth/refresh-token");

          const  accessToken = rs.data.Data;

	  useAuthStore.getState().setAccessToken(accessToken)

          return instance(originalConfig);
        } catch (_error) {
          return Promise.reject(_error);
        }
      }
    }

    return Promise.reject(err);
  }
);

export default instance