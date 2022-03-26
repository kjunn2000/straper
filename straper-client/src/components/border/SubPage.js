import { FaWindowClose } from "react-icons/fa";
import { useHistory } from "react-router-dom";
import { darkGrayBg } from "../../utils/style/color";
import { AiFillHome } from "react-icons/ai";
import { useRef } from "react";

function SubPage({ children }) {
  const history = useHistory();
  const sideBar = useRef();

  return (
    <div className="w-full min-h-screen flex flex-col lg:flex-row">
      {/* Mobile View */}
      <div
        className="sticky w-full top-0 text-gray-100 flex justify-between lg:hidden"
        style={{ background: "rgb(32,34,37)" }}
      >
        <a className="block p-4 text-white font-bold skew-x-3 skew-y-3">
          STRAPER
        </a>
        <div>
          <button
            className="mobile-menu-button p-4 focus:outline-none focus:bg-gray-700 hover:bg-indigo-600 transition duration-150"
            onClick={() => history.push("/channel")}
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
          onClick={() => history.push("/channel")}
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
