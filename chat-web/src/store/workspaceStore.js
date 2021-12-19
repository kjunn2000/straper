import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useWorkspaceStore = create((set) => ({
  workspaces: getLocalStorage("workspaces") || [],
  currWorkspace: getLocalStorage("currWorkspace") || {},
  currChannel: getLocalStorage("currChannel") || {},
  selectedChannelIds: getLocalStorage("selectedChannelIds") || new Map(),
  setWorkspaces: (workspaces) => {
    set((state) => {
      setLocalStorage("workspaces", workspaces);
      return { workspaces };
    });
  },
  addWorkspace: (workspace) => {
    set((state) => {
      const newWorkspaces = [...state.workspaces, workspace];
      setLocalStorage("workspaces", newWorkspaces);
      return { workspaces: newWorkspaces };
    });
  },
  deleteWorkspace: (workspaceId) => {
    set((state) => {
      const newWorkspaces = state.workspaces.filter(
        (workspace) => workspace.workspace_id != workspaceId
      );
      setLocalStorage("workspaces", newWorkspaces);
      return { workspaces: newWorkspaces };
    });
  },
  setCurrWorkspace: (workspaceId) => {
    set((state) => {
      const currWorkspace = state.workspaces.find(
        (workspace) => workspace.workspace_id == workspaceId
      );
      setLocalStorage("currWorkspace", currWorkspace);
      return { currWorkspace: currWorkspace };
    });
  },
  resetCurrWorkspace: () => {
    set((state) => {
      let newCurrWorkspace =
        state.workspaces.length > 0 ? state.workspaces[0] : {};
      setLocalStorage("currWorkspace", newCurrWorkspace);
      return { currWorkspace: newCurrWorkspace };
    });
  },
  setCurrChannel: (channelId) => {
    set((state) => {
      const currChannel = state.currWorkspace?.channel_list.find(
        (channel) => channel.channel_id == channelId
      );
      setLocalStorage("currChannel", currChannel);
      return { currChannel: currChannel };
    });
  },
  resetCurrChannel: () => {
    set((state) => {
      const currChannel =
        state.currWorkspace?.channel_list.length > 0
          ? state.currWorkspace?.channel_list[0]
          : {};
      return { currChannels: currChannel };
    });
  },

  setSelectedChannelIds: (channelId) => {
    set((state) => {
      const workspaceId = state.currWorkspace?.workspaceId;
      const newSelectedChannelIds = state.selectedChannelIds;
      newSelectedChannelIds.set(workspaceId, channelId);
      return { selectedChannelIds: newSelectedChannelIds };
    });
  },

  clearWorkspaceState: () => {
    set((state) => ({
      workspaces: [],
      currWorkspace: {},
    }));
  },
}));

export default useWorkspaceStore;
