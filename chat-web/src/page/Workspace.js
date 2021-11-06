import { useEffect } from "react";
import { useParams } from "react-router";
import ChatRoom from "../components/ChatRoom";
import Sidebar from "../components/sideBar/Sidebar";
import useWorkspaceStore from "../store/workspaceStore";

function Workspace() {
  const { workspaceId, channelId } = useParams();
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const resetCurrWorkspace = useWorkspaceStore(
    (state) => state.resetCurrWorkspace
  );
  const resetCurrChannel = useWorkspaceStore((state) => state.resetCurrChannel);
  const selectedChannelIds = useWorkspaceStore(
    (state) => state.selectedChannelIds
  );
  const setSelectedChannelIds = useWorkspaceStore(
    (state) => state.setSelectedChannelIds
  );
  useEffect(() => {
    updateWorkspaceState();
  }, []);

  const updateWorkspaceState = () => {
    console.log(workspaceId);
    console.log(channelId);
    if (workspaceId && workspaceId != "") {
      console.log("Has workspace id");
      setCurrWorkspace(workspaceId);
    } else {
      console.log("No workspace id");
      resetCurrWorkspace();
      selectPrevChannel();
      return;
    }
    if (channelId && channelId != "") {
      console.log("Has channel id");
      setCurrChannel(channelId);
      setSelectedChannelIds(channelId);
    } else {
      console.log("No channel id");
      selectPrevChannel();
    }
  };

  const selectPrevChannel = (workspaceId) => {
    const selectedChannelId = selectedChannelIds.get(workspaceId);
    if (selectedChannelId) {
      setCurrChannel(selectedChannelId);
    } else {
      resetCurrChannel();
    }
  };

  return (
    <div className="flex flex-row">
      <Sidebar />
      <ChatRoom />
    </div>
  );
}

export default Workspace;
