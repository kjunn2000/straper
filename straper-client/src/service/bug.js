import useIssueStore from "../store/issueStore";
import useWorkspaceStore from "../store/workspaceStore";
import api from "../axios/api";

export const getIssueData = async () => {
  console.log("getting issue...");
  const currWorkspace = useWorkspaceStore.getState().currWorkspace;
  const res = await api.get(
    `/protected/issue/list/${currWorkspace.workspace_id}?limit=100&offset=0`
  );
  if (res.data.Success && res.data.Data) {
    useIssueStore.getState().setIssues(res.data.Data);
  }
};
