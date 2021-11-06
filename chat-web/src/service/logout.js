import useAuthStore from "../store/authStore";
import useIdentifyStore from "../store/identityStore";
import useWorkspaceStore from "../store/workspaceStore";
import history from "./history";

export const logOut = (timeout) => {
  window.localStorage.clear();
  useAuthStore.getState()?.clearAccessToken();
  useIdentifyStore.getState()?.clearIdentity();
  useWorkspaceStore.getState()?.clearWorkspaceState();
  history.push("/login" + timeout && "/timeout");
  window.location.reload();
};
