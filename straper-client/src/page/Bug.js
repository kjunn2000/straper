import React, { useState, useMemo, useEffect } from "react";
import SubPage from "../components/border/SubPage";
import CreateIssueDialog from "../components/bug/CreateIssueDialog";
import Table, {
  AvatarCell,
  DateCell,
  SelectColumnFilter,
  StatusPill,
  TypeCell,
} from "../components/bug/Table";
import api from "../axios/api";
import useWorkspaceStore from "../store/workspaceStore";
import useIssueStore from "../store/issueStore";

const Bug = () => {
  const columns = useMemo(
    () => [
      // {
      //   Header: "Name",
      //   accessor: "name",
      //   Cell: AvatarCell,
      //   imgAccessor: "imgUrl",
      //   emailAccessor: "email",
      // },
      {
        Header: "T",
        accessor: "type",
        Cell: TypeCell,
        Filter: SelectColumnFilter, // new
        filter: "includes",
      },
      {
        Header: "Summary",
        accessor: "summary",
      },
      {
        Header: "Assignee",
        accessor: "assignee",
      },
      {
        Header: "Reporter",
        accessor: "reporter",
      },
      {
        Header: "P",
        accessor: "backlog_priority",
      },
      {
        Header: "Status",
        accessor: "status",
        Cell: StatusPill,
      },
      {
        Header: "Due",
        accessor: "due_time",
        Cell: DateCell,
      },
      // {
      //   Header: "Role",
      //   accessor: "role",
      //   Filter: SelectColumnFilter, // new
      //   filter: "includes",
      // },
    ],
    []
  );

  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const setIssues = useIssueStore((state) => state.setIssues);
  const issues = useIssueStore((state) => state.issues);

  useEffect(() => {
    getData();
  }, []);

  const getData = async () => {
    const res = await api.get(
      `/protected/issue/list/${currWorkspace.workspace_id}?limit=100&offset=0`
    );
    if (res.data.Success && res.data.Data) {
      setIssues(res.data.Data);
    }
  };

  const [createIssueDialogOpen, setCreateIssueDialogOpen] = useState(false);

  return (
    <SubPage>
      <div className="overflow-x-auto">
        <div className="flex justify-between">
          <span className="font-semibold font-size-xl text-xl">Issues</span>
          <button
            className="bg-indigo-500 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded"
            onClick={() => {
              setCreateIssueDialogOpen(true);
            }}
          >
            Create
          </button>
        </div>
        <div className="mt-4">
          <Table columns={columns} data={issues} />
        </div>
      </div>
      <CreateIssueDialog
        isOpen={createIssueDialogOpen}
        closeDialog={() => setCreateIssueDialogOpen(false)}
      />
    </SubPage>
  );
};

export default Bug;
