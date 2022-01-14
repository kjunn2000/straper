import React from "react";
import WorkspaceSidebar from "./WorkspaceSidebar";
import WorkspaceNavigator from "./WorkspaceNavigator";

function Sidebar() {
  return (
    <div className="flex flex-row">
      <WorkspaceNavigator />
      <WorkspaceSidebar />
    </div>
  );
}

export default Sidebar;
