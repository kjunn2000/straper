import ChatRoom from "../components/chat/ChatRoom";
import Sidebar from "../components/sideBar/Sidebar";
import useWorkspaceStore from "../store/workspaceStore";
import { ReactComponent as Welcome } from "../asset/img/welcome.svg";
import { useEffect } from "react";
import {
  fetchWorkspaceAccountList,
  fetchWorkspaceData,
  redirectToLatestWorkspace,
} from "../service/workspace";
import { isEmpty } from "../service/object";
import { connect, isSocketOpen } from "../service/websocket";
import { darkGrayBg } from "../utils/style/color";
import UserList from "../components/sideBar/UserList";

function Workspace() {
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const addIntervalId = useWorkspaceStore((state) => state.addIntervalId);
  const clearIntervalIds = useWorkspaceStore((state) => state.clearIntervalIds);

  useEffect(() => {
    fetchWorkspaceData().then((data) => redirectToLatestWorkspace(data));
    if (!isSocketOpen()) {
      connect();
    }
  }, []);

  useEffect(() => {
    if (!currWorkspace || !currWorkspace.workspace_id) {
      return;
    }
    clearIntervalIds();
    fetchWorkspaceAccountList(currWorkspace.workspace_id);
    addIntervalId(
      setInterval(() => {
        fetchWorkspaceAccountList(currWorkspace.workspace_id);
      }, 60000)
    );
  }, [currWorkspace.workspace_id]);

  const emptyComponent = (Image, text) => (
    <div className="flex flex-col items-center p-5 text-white">
      <div>{text}</div>
      <Image className="w-80 h-80" />
    </div>
  );

  return (
    <div className="flex" style={{ background: "rgb(54,57,63)" }}>
      <div className="flex-none">
        <Sidebar />
      </div>
      <div className="flex-auto">
        {isEmpty(currWorkspace) || isEmpty(currChannel) ? (
          emptyComponent(
            Welcome,
            "WELCOME TO STRAPER, LET'S STRENGTHEN YOUR COLLABORATION"
          )
        ) : (
          <ChatRoom />
        )}
      </div>
      <div className="flex-auto" style={darkGrayBg}>
        <UserList />
      </div>
    </div>
  );
}

export default Workspace;
