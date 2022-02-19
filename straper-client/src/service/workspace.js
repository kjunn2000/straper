import api from "../axios/api";
import useWorkspaceStore from "../store/workspaceStore";
import history from "./history";

export const fetchWorkspaceData = async () => {
  const workspaceState = useWorkspaceStore.getState();
  const res = await api.get("/protected/workspace/list");
  if (res.data?.Success && res.data?.Data) {
    workspaceState.setWorkspaces(res.data?.Data);
    workspaceState.setDefaultSelectedChannelIds();
    const selectedIds = [...workspaceState.selectedChannelIds];
    if (selectedIds.length > 0) {
      workspaceState.setCurrWorkspace(selectedIds[0][0]);
      workspaceState.setCurrChannel(selectedIds[0][1]);
    } else {
      workspaceState.clearWorkspaceState();
    }
  }
  return res.data?.Data;
};

export const fetchWorkspaceAccountList = async (workspaceId) => {
  const res = await api.get(`/protected/account/list/${workspaceId}`);
  if (res.data?.Success && res.data?.Data) {
    const resData = res.data.Data;
    const result = resData.reduce((map, obj) => {
      map[obj.user_id] = obj;
      return map;
    }, {});
    useWorkspaceStore.getState().setCurrAccountList(result);
  }
};

export const redirectToLatestWorkspace = (workspaces) => {
  let redirectLink = "/channel";
  if (workspaces.length > 0) {
    redirectLink += "/" + workspaces[0].workspace_id;
    if (workspaces[0].channel_list.length > 0) {
      redirectLink += "/" + workspaces[0].channel_list[0].channel_id;
    }
  }
  history.push(redirectLink);
};
