import { useParams } from "react-router";
import ChatRoom from "../components/ChatRoom";
import Sidebar from "../components/sideBar/Sidebar";
import useWorkspaceStore from "../store/workspaceStore";
import {ReactComponent as Welcome} from "../asset/img/welcome.svg";
import {ReactComponent as NotFound} from "../asset/img/notfound.svg";
import { useEffect } from "react";
import { fetchWorkspaceData, redirectToLatestWorkspace } from "../service/workspace";

function Workspace() {
  const { workspaceId, channelId } = useParams();
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);

  useEffect(() => {
      fetchWorkspaceData().then(data => redirectToLatestWorkspace(data));
  },[])

  const emptyComponent = (Image, text) => (
    <div className="flex flex-col items-center p-5 text-white">
        <div>{text}</div>
        <Image className="w-80 h-80"/>
    </div>
  )

  return (
    <div className="flex" 
      style={{ background: "rgb(54,57,63)" }}
    >
      <div className="flex-none">
        <Sidebar/>
      </div>
      <div className="flex-auto">
      {
        !workspaceId ? emptyComponent(Welcome, "WELCOME TO STRAPER, LET'S STRENGTHEN YOUR COLLABORATION") 
        : (!currWorkspace || currWorkspace == {} ? emptyComponent(NotFound, "WORKSPACE NOT FOUND") 
        : <ChatRoom/>)
      }
      </div>
    </div>
  );
}

export default Workspace;
