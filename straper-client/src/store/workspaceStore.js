import create from "zustand";
import { getLocalStorage, setLocalStorage } from "./localStorage";

const useWorkspaceStore = create((set) => ({
  workspaces: getLocalStorage("workspaces") || [],
  currWorkspace: getLocalStorage("currWorkspace") || {},
  currChannel: getLocalStorage("currChannel") || {},
  currAccountList: getLocalStorage("currAccountList") || {},
  selectedChannelIds:
    new Map(getLocalStorage("selectedChannelIds")) || new Map(),
  intervalIds: getLocalStorage("intervalIds") || [],
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
      state.clearIntervalIds();
      const newWorkspaces = state.workspaces.filter(
        (workspace) => workspace.workspace_id !== workspaceId
      );
      setLocalStorage("workspaces", newWorkspaces);
      return { workspaces: newWorkspaces };
    });
  },
  setCurrWorkspace: (workspaceId) => {
    set((state) => {
      const currWorkspace = state.workspaces.find(
        (workspace) => workspace.workspace_id === workspaceId
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
  addChannelToWorkspace: (workspaceId, channel) => {
    set((state) => {
      const workspaces = state.workspaces.map((workspace) => {
        if (workspace.workspace_id === workspaceId) {
          workspace.channel_list.push(channel);
        }
        return workspace;
      });
      setLocalStorage("workspaces", workspaces);
      return { workspaces: workspaces };
    });
  },
  deleteChannelFromWorkspace: (channelId) => {
    set((state) => {
      const workspaces = state.workspaces.map((workspace) => {
        if (workspace.workspace_id === state.currWorkspace.workspace_id) {
          workspace.channel_list = workspace.channel_list.filter(
            (channel) => channel.channel_id !== channelId
          );
        }
        return workspace;
      });
      setLocalStorage("workspaces", workspaces);
      return { workspaces: workspaces };
    });
  },
  setCurrChannel: (channelId) => {
    set((state) => {
      const currChannel = state.currWorkspace?.channel_list.find(
        (channel) => channel.channel_id === channelId
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
      setLocalStorage("currChannel", currChannel);
      return { currChannels: currChannel };
    });
  },
  setDefaultSelectedChannelIds: () => {
    set((state) => {
      const newSelectedChannelIds = state.selectedChannelIds;
      state.workspaces.forEach((workspace) => {
        if (workspace.channel_list.length > 0) {
          newSelectedChannelIds.set(
            workspace.workspace_id,
            workspace.channel_list[0].channel_id
          );
        }
      });
      setLocalStorage(
        "selectedChannelIds",
        Array.from(newSelectedChannelIds.entries)
      );
      return { selectedChannelIds: newSelectedChannelIds };
    });
  },
  setSelectedChannelIds: (workspaceId, channelId) => {
    set((state) => {
      const newSelectedChannelIds = state.selectedChannelIds;
      newSelectedChannelIds.set(workspaceId, channelId);
      setLocalStorage(
        "selectedChannelIds",
        Array.from(newSelectedChannelIds.entries)
      );
      return { selectedChannelIds: newSelectedChannelIds };
    });
  },
  deleteSelectedChannelIds: (workspaceId) => {
    set((state) => {
      const newSelectedChannelIds = state.selectedChannelIds;
      newSelectedChannelIds.delete(workspaceId);
      setLocalStorage(
        "selectedChannelIds",
        Array.from(newSelectedChannelIds.entries)
      );
      return { selectedChannelIds: newSelectedChannelIds };
    });
  },

  clearWorkspaceState: () => {
    set((state) => ({
      workspaces: [],
      currWorkspace: {},
      currChannel: {},
      selectedChannelIds: new Map(),
      intervalIds: [],
      currAccountList: {},
    }));
    setLocalStorage("workspaces", []);
    setLocalStorage("currWorkspace", {});
    setLocalStorage("currChannel", {});
    setLocalStorage("selectedChannelIds", []);
    setLocalStorage("currAccountList", {});
    setLocalStorage("intervalIds", []);
  },

  setCurrAccountList: (accountList) => {
    setLocalStorage("currAccountList", accountList);
    set((state) => ({
      currAccountList: accountList,
    }));
  },

  addIntervalId: (id) => {
    set((state) => {
      if (!state.intervalIds) {
        state.intervalIds = [];
      }
      const newIntervalIds = [...state.intervalIds, id];
      setLocalStorage("intervalIds", newIntervalIds);
      return { intervalIds: newIntervalIds };
    });
  },

  clearIntervalIds: () => {
    set((state) => {
      state.intervalIds.forEach((id) => clearInterval(id));
      setLocalStorage("intervalIds", []);
      return { intervalIds: [] };
    });
  },
}));

export default useWorkspaceStore;
