import { FaWindowClose } from "react-icons/fa";
import { useHistory } from "react-router-dom";
import { darkGrayBg } from "../../utils/style/color";

function SubPage({ children }) {
  const history = useHistory();

  return (
    <div className="w-full h-screen grid grid-cols-10">
      <div
        className="p-3 bg-gray-200"
        // style={darkGrayBg}
      >
        <FaWindowClose
          size="40"
          className="text-indigo-500 cursor-pointer"
          onClick={() => history.push("/channel")}
        />
      </div>
      <div
        className="col-span-9 flex flex-col p-5"
        // style={{ background: "rgb(54,57,63)" }}
      >
        {children}
      </div>
    </div>
  );
}

export default SubPage;
