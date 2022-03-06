import { useEffect } from "react";
import api from "../axios/api";
import useWorkspaceStore from "../store/workspaceStore";
import useBoardStore from "../store/boardStore";
import { isEmpty } from "../service/object";
import DragList from "../components/board/DragList";
import { connect, isSocketOpen } from "../service/websocket";
import SubPage from "../components/border/SubPage";
import { fetchWorkspaceAccountList } from "../service/workspace";

function TaskBoard() {
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
    fetchWorkspaceAccountList(currWorkspace.workspace_id);
    if (!isSocketOpen()) {
      connect();
    }
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

  return (
    <SubPage>
      <span className="text-gray-500 font-bold text-center">
        {board.board_name}
      </span>
      <DragList />
    </SubPage>
  );
}

export default TaskBoard;
