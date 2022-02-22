import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useIssueStore = create((set) => ({
  issues: getLocalStorage("issues") || [],
  assigneeOptions: getLocalStorage("assigneeOptions") || {},
  setIssues: (issues) => {
    setLocalStorage("issues", issues);
    set(() => ({
      issues: issues,
    }));
  },
  clearIssues: () => {
    setLocalStorage("issues", []);
    set(() => ({
      issues: [],
    }));
  },
  addIssue: (issue) => {
    set((state) => {
      const newIssues = [...state.issues, issue];
      setLocalStorage("issues", newIssues);
      return {
        issues: newIssues,
      };
    });
  },
  updateIssue: (issue) => {
    set((state) => {
      const newIssues = state.issues.map((i) =>
        i.issue_id === issue.issue_id ? issue : i
      );
      setLocalStorage("issues", newIssues);
      return {
        issues: newIssues,
      };
    });
  },
  deleteIssue: (issueId) => {
    set((state) => {
      const newIssues = state.issues.filter(
        (issue) => issue.issue_id !== issueId
      );
      setLocalStorage("issues", newIssues);
      return {
        issues: newIssues,
      };
    });
  },
  setAssigneeOptions: (options) => {
    const map = options.reduce((map, obj) => {
      map[obj.user_id] = obj.username;
      return map;
    }, {});
    set(() => ({ assigneeOptions: map }));
  },
}));

export default useIssueStore;
