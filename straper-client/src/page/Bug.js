import React, { useState, useMemo, useEffect } from "react";
import SubPage from "../components/border/SubPage";
import Table, {
  DateCell,
  SelectColumnFilter,
  StatusPill,
  SummaryCell,
  TypeCell,
} from "../components/bug/Table";
import useIssueStore from "../store/issueStore";
import IssueDialog from "../components/bug/IssueDialog";
import { getIssueData } from "../service/bug";

const Bug = () => {
  const columns = useMemo(
    () => [
      {
        Header: "Type",
        accessor: "type",
        Cell: TypeCell,
        Filter: SelectColumnFilter,
        filter: "includes",
      },
      {
        Header: "Summary",
        accessor: "summary",
        Cell: SummaryCell,
        idAccessor: "issue_id",
      },
      {
        Header: "Assignee",
        accessor: "assignee_name",
      },
      {
        Header: "Reporter",
        accessor: "reporter_name",
      },
      {
        Header: "Priority",
        accessor: "backlog_priority",
      },
      {
        Header: "Status",
        accessor: "status",
        Cell: StatusPill,
      },
      {
        Header: "Due Time",
        accessor: "due_time",
        Cell: DateCell,
      },
    ],
    []
  );

  const issues = useIssueStore((state) => state.issues);

  useEffect(() => {
    getIssueData();
  }, []);

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
      <IssueDialog
        isOpen={createIssueDialogOpen}
        closeDialog={() => setCreateIssueDialogOpen(false)}
      />
    </SubPage>
  );
};

export default Bug;
