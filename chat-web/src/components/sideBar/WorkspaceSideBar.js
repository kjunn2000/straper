import React, { Fragment, useState } from "react";
import SidebarIcon from "../SidebarIcon";
import useWorkspaceStore from "../../store/workspaceStore";
import AddWorkspaceDialog from "../dialog/AddWorkspaceDialog";
import JoinWorkspaceDialog from "../dialog/JoinWorkspaceDialog";
import useAuthStore from "../../store/authStore";
import useIdentifyStore from "../../store/identityStore";
import { useHistory } from "react-router";
import { logOut } from "../../service/logout";

function WorkspaceSidebar() {
  const workspaces = useWorkspaceStore((state) => state.workspaces);

  const clearToken = useAuthStore((state) => state.clearAccessToken);
  const clearIdentity = useIdentifyStore((state) => state.clearIdentity);
  const clearWorkspaceState = useWorkspaceStore(
    (state) => state.clearWorkspaceState
  );

  const history = useHistory();

  const [isAddWorkspaceDialogOpen, setAddWorkspaceDialogOpen] = useState(false);
  const [isJoinWorkspaceDialogOpen, setJoinWorkspaceDialogOpen] =
    useState(false);

  const changeWorkspace = (workspaceId) => {
    history.push(`/channels/${workspaceId}`);
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
      <AddWorkspaceDialog
        isOpen={isAddWorkspaceDialogOpen}
        close={() => setAddWorkspaceDialogOpen(false)}
        toggleDialog={toggleDialog}
      />
      <JoinWorkspaceDialog
        isOpen={isJoinWorkspaceDialogOpen}
        close={() => setJoinWorkspaceDialogOpen(false)}
        toggleDialog={toggleDialog}
      ></JoinWorkspaceDialog>
    </div>
  );
}

export default WorkspaceSidebar;
