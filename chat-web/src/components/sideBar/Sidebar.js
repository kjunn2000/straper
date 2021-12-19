import React from "react";
import WorkspaceSidebar from "./WorkspaceSideBar";
import ChannelSidebar from "./ChannelSidebar";
import useWorkspaceStore from "../../store/workspaceStore";

function Sidebar() {

  const workspace = useWorkspaceStore((state) => state.currWorkspace);

  return (
    <div className="flex flex-row">
      <WorkspaceSidebar />
      { workspace == {} ? 
        <ChannelSidebar />
        :  <></>
      }
    </div>
  );
}

export default Sidebar;
