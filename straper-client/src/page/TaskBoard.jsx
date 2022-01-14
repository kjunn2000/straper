import { darkGrayBg } from "../utils/style/color";
import { FaWindowClose } from "react-icons/fa";
import { useHistory } from "react-router-dom";

function TaskBoard() {

    const history = useHistory()

  return (
    <div className="w-full h-screen grid grid-cols-10" style={darkGrayBg}>
       <div className="p-3">
            <FaWindowClose
              size="40"
              className="text-indigo-500 cursor-pointer"
              onClick={() => history.push("/channel")}
            />
           </div> 
       <div className="col-span-9 bg-sky-200">
           
           </div> 
    </div>
  );
}

export default TaskBoard;
