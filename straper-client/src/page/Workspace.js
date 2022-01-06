import { useParams } from "react-router";
import ChatRoom from "../components/chat/ChatRoom";
import Sidebar from "../components/sideBar/Sidebar";
import useWorkspaceStore from "../store/workspaceStore";
import { ReactComponent as Welcome } from "../asset/img/welcome.svg";
import { ReactComponent as NotFound } from "../asset/img/notfound.svg";
import { useEffect } from "react";
import {
  fetchWorkspaceData,
  redirectToLatestWorkspace,
} from "../service/workspace";
import { isEmpty } from "../service/object";
import { connect } from "../service/websocket";
import useMessageStore from "../store/messageStore";

function Workspace() {
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const pushMessage = useMessageStore((state) => state.pushMessage);

  useEffect(() => {
    fetchWorkspaceData().then((data) => redirectToLatestWorkspace(data));
    connect(pushMessage);
  }, []);

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
    </div>
  );
}

export default Workspace;
