import React, { useState } from "react";
import SidebarIcon from "../SidebarIcon";
import useWorkspaceStore from "../../store/workspaceStore";
import { useHistory } from "react-router";
import { logOut } from "../../service/logout";
import AddDialog from "../dialog/AddDialog";
import api from "../../axios/api";
import JoinDialog from "../dialog/JoinDialog";

function WorkspaceSidebar() {
  const workspaces = useWorkspaceStore((state) => state.workspaces);
  const selectedChannelIds = useWorkspaceStore((state) => state.selectedChannelIds);
  const addWorkspace = useWorkspaceStore((state) => state.addWorkspace);
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const setSelectedChannelIds = useWorkspaceStore((state) => state.setSelectedChannelIds);

  const history = useHistory();

  const [isAddWorkspaceDialogOpen, setAddWorkspaceDialogOpen] = useState(false);
  const [isJoinWorkspaceDialogOpen, setJoinWorkspaceDialogOpen] =
    useState(false);
  const [errMsg, setErrMsg] = useState("");

  const changeWorkspace = (workspaceId) => {
    const channelId = selectedChannelIds.get(workspaceId);
    setCurrWorkspace(workspaceId);
    setCurrChannel(channelId);
    history.push(`/channel/${workspaceId}/${channelId}`);
  };

  const toggleDialog = () => {
    if (isAddWorkspaceDialogOpen) {
      setAddWorkspaceDialogOpen(false);
      setJoinWorkspaceDialogOpen(true);
    } else {
      setJoinWorkspaceDialogOpen(false);
      setAddWorkspaceDialogOpen(true);
    }
  };

  const addNewWorkspace = (data) => {
    api.post("/protected/workspace/create", data).then((res) => {
      if (res.data.Success) {
        updateNewWorkspace(res.data.Data);
      }
    });
  };

  const joinNewWorkspace = async (data) => {
    const res = await api.post(`/protected/workspace/join/${data?.workspace_id}`)
    if (res.data.Success) {
      updateNewWorkspace(res.data.Data);
      return;
    }else {
      switch (res.data.ErrorMessage) {
        case "workspace.user.record.exist": {
          return "You has been joined to this workspace."
        }
        case "invalid.workspace.id": {
          return "Invalid workspace id."
        }
      }
    }
  };

  const updateNewWorkspace = (newWorkspace) => {
    addWorkspace(newWorkspace);
    setCurrWorkspace(newWorkspace.workspace_id)
    const channelId = newWorkspace.channel_list[0].channel_id
    setCurrChannel(channelId)
    setSelectedChannelIds(newWorkspace.workspace_id, channelId)
    history.push(`/channel/${newWorkspace.workspace_id}/${channelId}`)
  }

  return (
    <div>
      <div
        className="flex flex-col w-24 h-screen p-3"
        style={{ background: "rgb(32,34,37)" }}
      >
        {workspaces &&
          workspaces.map((workspace) => (
            <SidebarIcon
              key={workspace.workspace_id}
              content={workspace.workspace_name}
              click={() => changeWorkspace(workspace.workspace_id)}
              bgColor="bg-gray-500"
              hoverBgColor="hover:bg-indigo-500"
            />
          ))}
        <SidebarIcon
          content="+"
          click={() => setAddWorkspaceDialogOpen(true)}
          bgColor="bg-indigo-500"
          hoverBgColor="hover:bg-indigo-800"
        />
        <div className="absolute bottom-5 left-5">
          <SidebarIcon
            content="LogOut"
            click={() => logOut(false)}
            bgColor="bg-indigo-500"
            hoverBgColor="hover:bg-indigo-800"
          />
        </div>
      </div>
      <AddDialog
        isOpen={isAddWorkspaceDialogOpen}
        close={() => setAddWorkspaceDialogOpen(false)}
        toggleDialog={toggleDialog}
        addAction={addNewWorkspace}
        type="workspace"
      />
      <JoinDialog
        isOpen={isJoinWorkspaceDialogOpen}
        close={() => setJoinWorkspaceDialogOpen(false)}
        toggleDialog={toggleDialog}
        joinAction={joinNewWorkspace}
        type="workspace"
        errMsg={errMsg}
      />
    </div>
  );
}

export default WorkspaceSidebar;
