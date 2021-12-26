import React, { useState } from "react";
import { AiOutlinePlus } from "react-icons/ai";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import api from "../../axios/api";
import useWorkspaceStore from "../../store/workspaceStore";
import ActionDialog from "../dialog/ActionDialog";
import AddChannelDialog from "../dialog/AddChannelDialog";
import SimpleDialog from "../dialog/SimpleDialog";
import WorkspaceMenu from "../WorkspaceMenu";

function ChannelSidebar() {
  const workspace = useWorkspaceStore((state) => state.currWorkspace);
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const setSelectedChannelIds = useWorkspaceStore((state) => state.setSelectedChannelIds);
  const deleteChannelFromWorkspace = useWorkspaceStore((state) => state.deleteChannelFromWorkspace);
  const selectedChannelIds = useWorkspaceStore((state) => state.selectedChannelIds);
  const [addChannelDialogOpen, setAddChannelDialogOpen] = useState(false);
  const history = useHistory();
  const [failDeleteDialogOpen, setFailDeleteDialogOpen] = useState(false);
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [targetDeleteChannelId, setTargetDeleteChannelId] = useState("");

  const changeChannel = (channelId) => {
    setCurrChannel(channelId);
    setSelectedChannelIds(workspace.workspace_id, channelId)
    history.push(`/channel/${workspace.workspace_id}/${channelId}`)
  }

  const onDeleteChannel = (channelId) => {
    if (workspace.channel_list.length == 1) {
      setFailDeleteDialogOpen(true);
      return
    }
    setTargetDeleteChannelId(channelId);
    setDeleteWarningDialogOpen(true);
  }

  const deleteChannel = async (channelId) => {
    const res = await api.post(`/protected/channel/delete/${channelId}`);
    if (res.data.Success){
      deleteChannelFromWorkspace(channelId);
      setCurrWorkspace(workspace.workspace_id);
      if (selectedChannelIds.get(workspace.workspace_id) == channelId) {
        const nextChannelId = workspace.channel_list[0].channel_id
        setCurrChannel(nextChannelId);   
        setSelectedChannelIds(workspace.workspace_id, nextChannelId)  
        history.push(`/channel/${workspace.workspace_id}/${nextChannelId}`)
      }
    }
  }

  return (
    workspace.workspace_id ? 
    <div
      className="flex flex-col w-64 h-screen "
      style={{ background: "rgb(47,49,54)" }}
    >
      <WorkspaceMenu/>
      <div className="p-5 text-sm text-gray-400 flex justify-between hover:text-white">
              <span>CHANNELS</span>
              <AiOutlinePlus onClick={()=>setAddChannelDialogOpen(true)}/>
        </div>
      <div className="px-3">
        {workspace?.channel_list &&
          workspace.channel_list.map((channel) => (
            <div
              className="group flex justify-between text-white text-sm font-medium p-3 text-gray-400 hover:bg-gray-700 rounded hover:text-white"
              key={channel?.channel_id}
              onClick={()=>changeChannel(channel.channel_id)}
            >
              <span> # {channel?.channel_name} </span>
              <span className="opacity-0 group-hover:opacity-100 cursor-pointer"
                onClick={()=>onDeleteChannel(channel.channel_id)}
              >x</span>
            </div>
          ))}
      </div>
      <AddChannelDialog isOpen={addChannelDialogOpen} close={()=>setAddChannelDialogOpen(false)}/>
      <SimpleDialog
        isOpen={failDeleteDialogOpen}
        setIsOpen={setFailDeleteDialogOpen}
        title="Fail To Delete Last Channel"
        content="Unfortunately to tell you that one workspace should has at least one channel."
        buttonText="Close"
        buttonStatus="fail"
      />
      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Delete Channel Confirmation"
        content="Please confirm that the deleted channel will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={()=>deleteChannel(targetDeleteChannelId)}
        closeButtonText="close"
      />
    </div>
      :
    <div
      className="flex flex-col w-64 h-screen "
      style={{ background: "rgb(47,49,54)" }}
    ></div>
  );
}

export default ChannelSidebar;
