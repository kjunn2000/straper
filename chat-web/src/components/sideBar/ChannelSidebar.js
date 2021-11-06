import React from "react";
import useWorkspaceStore from "../../store/workspaceStore";
import WorkspaceMenu from "../WorkspaceMenu";

function ChannelSidebar() {
  const workspace = useWorkspaceStore((state) => state.currWorkspace);

  return (
    <div
      className="flex flex-col w-64 h-screen "
      style={{ background: "rgb(47,49,54)" }}
    >
      <WorkspaceMenu />
      <div className="px-3">
        {workspace?.channel_list &&
          workspace.channel_list.map((channel) => (
            <div
              className="text-white font-medium p-3 hover:bg-gray-700 rounded-lg"
              key={channel?.channel_id}
            >
              {" "}
              # {channel?.channel_name}
            </div>
          ))}
      </div>
    </div>
  );
}

export default ChannelSidebar;
