import useAuthStore from "../store/authStore";
import useIdentityStore from "../store/identityStore";
import history from "./history";

export const logOut = (timeout) => {
  window.localStorage.clear();
  useAuthStore.getState()?.clearAccessToken();
  useIdentityStore.getState()?.clearIdentity();
  const url = "/login" + (timeout ? "/timeout" : "");
  history.push(url);
};
