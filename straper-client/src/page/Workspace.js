import ChatRoom from "../components/chat/ChatRoom";
import Sidebar from "../components/sideBar/Sidebar";
import useWorkspaceStore from "../store/workspaceStore";
import { ReactComponent as Welcome } from "../asset/img/welcome.svg";
import { useEffect, useRef } from "react";
import {
  fetchWorkspaceAccountList,
  fetchWorkspaceData,
  redirectToLatestWorkspace,
} from "../service/workspace";
import { isEmpty } from "../service/object";
import { connect, isSocketOpen } from "../service/websocket";
import { darkGrayBg } from "../utils/style/color";
import UserList from "../components/sideBar/UserList";
import { MdOutlineWork } from "react-icons/md";
import { FaUsers } from "react-icons/fa";
import { AiFillCloseCircle } from "react-icons/ai";

function Workspace() {
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const addIntervalId = useWorkspaceStore((state) => state.addIntervalId);
  const clearIntervalIds = useWorkspaceStore((state) => state.clearIntervalIds);
  const workspaceSideBar = useRef();
  const usersSideBar = useRef();

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

  const toggleWorkspaceSideBar = () => {
    if (usersSideBar.current.style.display === "flex") {
      toggleUsersSideBar();
    }
    const display = workspaceSideBar.current.style.display === "" ? "flex" : "";
    workspaceSideBar.current.style.display = display;
  };

  const toggleUsersSideBar = () => {
    if (workspaceSideBar.current.style.display === "flex") {
      toggleWorkspaceSideBar();
    }
    const display = usersSideBar.current.style.display === "" ? "flex" : "";
    usersSideBar.current.style.display = display;
  };

  return (
    <div
      className="relative min-h-screen lg:flex"
      style={{ background: "rgb(54,57,63)" }}
    >
      {/* Mobile View */}
      <div className="absolute top-2 right-2 text-gray-100 flex space-x-3 lg:hidden">
        <button
          className="mobile-menu-button p-4 focus:outline-none bg-slate-800 focus:bg-gray-700 hover:bg-indigo-600 transition duration-150 rounded-full"
          onClick={() => toggleWorkspaceSideBar()}
        >
          <MdOutlineWork size={20} />
        </button>
        <button
          className="mobile-menu-button p-4 focus:outline-none bg-slate-800 focus:bg-gray-700 hover:bg-indigo-600 transition duration-150 rounded-full"
          onClick={() => toggleUsersSideBar()}
        >
          <FaUsers size={20} />
        </button>
      </div>
      <div
        className="w-1/5 absolute lg:relative inset-y-0 left-0 hidden lg:flex z-10"
        ref={workspaceSideBar}
      >
        <Sidebar />
      </div>
      <div className="w-full lg:w-3/5">
        {isEmpty(currWorkspace) || isEmpty(currChannel) ? (
          emptyComponent(
            Welcome,
            "WELCOME TO STRAPER, LET'S STRENGTHEN YOUR COLLABORATION"
          )
        ) : (
          <ChatRoom />
        )}
      </div>
      <div
        className="w-2/5 lg:w-1/5 absolute lg:relative inset-y-0 right-0 hidden flex-col lg:flex z-10"
        style={darkGrayBg}
        ref={usersSideBar}
      >
        <div className="lg:hidden flex justify-end p-2">
          <AiFillCloseCircle
            size={30}
            className="text-indigo-300"
            onClick={() => toggleUsersSideBar()}
          />
        </div>
        <UserList />
      </div>
    </div>
  );
}

export default Workspace;
