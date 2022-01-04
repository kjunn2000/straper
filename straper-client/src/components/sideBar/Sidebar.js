import React from "react";
import WorkspaceSidebar from "./WorkspaceSideBar";
import ChannelSidebar from "./ChannelSidebar";
import useWorkspaceStore from "../../store/workspaceStore";

function Sidebar() {

  return (
    <div className="flex flex-row">
      <WorkspaceSidebar />
      <ChannelSidebar />
    </div>
  );
}

export default Sidebar;
