import React, { useState, useMemo } from "react";
import SubPage from "../components/border/SubPage";
import CreateIssueDialog from "../components/bug/CreateIssueDialog";
import Table, {
  AvatarCell,
  SelectColumnFilter,
  StatusPill,
} from "../components/bug/Table";

const getData = () => {
  const data = [
    {
      type: "bug",
      summary: "fix the ui position",
      assignee: "Tam Kai Jun",
      reporter: "Chai Juo Ann",
      priority: "High",
      status: "Active",
      due_time: "22/02/22",
      created_date: "22/02/22",
    },
    {
      type: "epic",
      summary: "fix the ui position",
      assignee: "King Kong",
      reporter: "Chai Juo Ann",
      priority: "High",
      status: "Low",
      due_time: "22/02/22",
      created_date: "22/02/22",
    },
  ];
  return [...data, ...data, ...data];
};

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
        accessor: "priority",
      },
      {
        Header: "Status",
        accessor: "status",
        Cell: StatusPill,
      },
      {
        Header: "Due",
        accessor: "due_time",
      },
      {
        Header: "Created",
        accessor: "created_date",
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

  const data = useMemo(() => getData(), []);
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
          <Table columns={columns} data={data} />
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
