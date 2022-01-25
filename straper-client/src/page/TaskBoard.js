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
  const setTaskListsOrder = useBoardStore((state) => state.setTaskListsOrder);
  const setCurrAccountList = useWorkspaceStore(
    (state) => state.setCurrAccountList
  );

  useEffect(() => {
    getBoardData();
    getWorkspaceUsersInfo();
    connect();
  }, []);

  const getBoardData = () => {
    api.get(`/protected/board/${currWorkspace.workspace_id}`).then((res) => {
      if (res.data.Success) {
        const data = res.data.Data;
        setBoard(data.task_board);
        if (!isEmpty(data.task_lists)) {
          const result = data.task_lists.reduce((map, obj) => {
            if (!obj.card_list) {
              obj.card_list = [];
            }
            obj["card_list_order"] = obj.card_list.map((card) => card.card_id);
            obj["card_list"] = obj.card_list.reduce(
              (map, obj) => ((map[obj.card_id] = obj), map),
              {}
            );
            map[obj.list_id] = obj;
            return map;
          }, {});

          setTaskLists(result);
          const taskListsOrder = data.task_lists.map(
            (taskList) => taskList.list_id
          );
          setTaskListsOrder(taskListsOrder);
        }
      }
    });
  };

  const getWorkspaceUsersInfo = () => {
    api
      .get(`/protected/account/list/${currWorkspace.workspace_id}`)
      .then((res) => {
        if (res.data.Success) {
          const data = res.data.Data;
          if (!isEmpty(data)) {
            const userListMap = data.reduce((map, obj) => {
              map[obj.user_id] = obj;
              return map;
            }, {});
            setCurrAccountList(userListMap);
          }
        }
      });
  };

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
