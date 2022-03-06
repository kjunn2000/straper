import useAuthStore from "../store/authStore";
import useIdentityStore from "../store/identityStore";
import useIssueStore from "../store/issueStore";
import useWorkspaceStore from "../store/workspaceStore";
import history from "./history";
import { sendUnregisterMsg } from "./websocket";

export const logOut = (timeout) => {
  sendUnregisterMsg();
  window.localStorage.clear();
  useAuthStore.getState()?.clearAccessToken();
  useIdentityStore.getState()?.clearIdentity();
  useWorkspaceStore.getState()?.clearIntervalIds();
  useWorkspaceStore.getState()?.clearWorkspaceState();
  useIssueStore.getState()?.clearIssues();
  const url = "/login" + (timeout ? "/timeout" : "");
  history.push(url);
};
