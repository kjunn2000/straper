import { FaWindowClose } from "react-icons/fa";
import { useHistory } from "react-router-dom";
import { darkGrayBg } from "../../utils/style/color";
import { AiFillHome } from "react-icons/ai";
import { useRef } from "react";
import useWorkspaceStore from "../../store/workspaceStore";

function SubPage({ children }) {
  const history = useHistory();
  const sideBar = useRef();
  const workspace = useWorkspaceStore((state) => state.currWorkspace);
  const channel = useWorkspaceStore((state) => state.currChannel);

  const closeSubPage = () => {
    history.push(`/channel/${workspace.workspace_id}/${channel.channel_id}`);
  };

  return (
    <div className="w-full min-h-screen flex flex-col lg:flex-row">
      {/* Mobile View */}
      <div className="absolute w-full top-2 right-2 text-gray-100 flex justify-end lg:hidden">
        <div>
          <button
            className="mobile-menu-button p-4 focus:outline-none bg-slate-800 focus:bg-gray-700 hover:bg-indigo-600 transition duration-150 rounded-full"
            onClick={() => closeSubPage()}
          >
            <AiFillHome size={20} />
          </button>
        </div>
      </div>
      <div
        className="w-1/12 p-3 bg-gray-200 hidden lg:flex"
        style={darkGrayBg}
        ref={sideBar}
      >
        <FaWindowClose
          size="40"
          className="text-indigo-500 cursor-pointer"
          onClick={() => closeSubPage()}
        />
      </div>
      <div
        className="w-full min-h-screen lg:w-11/12 flex flex-col p-5"
        style={{ background: "rgb(54,57,63)" }}
      >
        {children}
      </div>
    </div>
  );
}

export default SubPage;
