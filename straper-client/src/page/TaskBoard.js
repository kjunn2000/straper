import { darkGrayBg } from "../utils/style/color";
import { FaWindowClose } from "react-icons/fa";
import { useHistory } from "react-router-dom";
import { useEffect } from "react";
import api from "../axios/api";
import useWorkspaceStore from "../store/workspaceStore";
import useBoardStore from "../store/boardStore";
import { isEmpty } from "../service/object";
import DragList from "../components/board/DragList";
import { connect } from "../service/websocket";

function TaskBoard() {
  const history = useHistory();
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const board = useBoardStore((state) => state.board);
  const setBoard = useBoardStore((state) => state.setBoard);
  const setTaskLists = useBoardStore((state) => state.setTaskLists);

  useEffect(() => {
    api.get(`/protected/board/${currWorkspace.workspace_id}`).then((res) => {
      if (res.data.Success) {
        const data = res.data.Data;
        setBoard(data.task_board);
        if (!isEmpty(data.task_lists)) {
          setTaskLists(data.task_lists);
        }
      }
    });
    connect();
  }, []);

  return (
    <div className="w-full h-screen grid grid-cols-10" style={darkGrayBg}>
      <div className="p-3">
        <FaWindowClose
          size="40"
          className="text-indigo-500 cursor-pointer"
          onClick={() => history.push("/channel")}
        />
      </div>
      <div
        className="col-span-9 flex flex-col overflow-x-auto p-5"
        style={{ background: "rgb(54,57,63)" }}
      >
        <span className="text-white font-bold text-center">
          {board.board_name}
        </span>
        <DragList />
      </div>
    </div>
  );
}

export default TaskBoard;
