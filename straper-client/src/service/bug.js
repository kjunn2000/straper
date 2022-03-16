import useIssueStore from "../store/issueStore";
import useWorkspaceStore from "../store/workspaceStore";
import api from "../axios/api";

export const getIssueData = async () => {
  const currWorkspace = useWorkspaceStore.getState().currWorkspace;
  const res = await api.get(
    `/protected/issue/list/${currWorkspace.workspace_id}?limit=100&offset=0`
  );
  if (res.data.Success && res.data.Data) {
    const currAccountList = useWorkspaceStore.getState().currAccountList;
    const newData = res.data.Data.map((issue) =>
      setAssigneeReporterName(issue, currAccountList)
    );
    useIssueStore.getState().setIssues(newData);
  }
};

export const setAssigneeReporterName = (issue, currAccountList) => {
  const assignee = currAccountList[issue.assignee];
  const reporter = currAccountList[issue.reporter];
  issue.assignee_name =
    assignee && assignee !== undefined ? assignee.username : "-";
  issue.reporter_name =
    reporter && reporter !== undefined ? reporter.username : "-";
  return issue;
};
