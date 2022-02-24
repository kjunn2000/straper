import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useIssueStore = create((set) => ({
  issues: getLocalStorage("issues") || [],
  isSet: getLocalStorage("isSet") || false,
  assigneeOptions: getLocalStorage("assigneeOptions") || {},
  setIssues: (issues) => {
    setLocalStorage("issues", issues);
    set(() => ({
      issues: issues,
      isSet: true,
    }));
  },
  clearIssues: () => {
    setLocalStorage("issues", []);
    set(() => ({
      issues: [],
      isSet: false,
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
  addIssueAttachments: (issueId, attachments) => {
    set((state) => {
      const newIssues = state.issues.map((issue) => {
        if (issue.issue_id === issueId) {
          if (!issue.attachments) {
            issue.attachments = [];
          }
          issue.attachments = [...issue.attachments, ...attachments];
        }
        return issue;
      });
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
  deleteAttachment: (issueId, fid) => {
    set((state) => {
      const newIssues = state.issues.map((issue) => {
        if (issue.issue_id === issueId) {
          issue.attachments = issue.attachments.filter((a) => a.fid !== fid);
        }
        return issue;
      });
      setLocalStorage("issues", newIssues);
      return {
        issues: newIssues,
      };
    });
  },
}));

export default useIssueStore;
