import React, { useMemo } from "react";
import Table, { SelectColumnFilter } from "../shared/table/Table";

const ManageUser = () => {
  const columns = useMemo(
    () => [
      {
        Header: "T",
        accessor: "type",
        Filter: SelectColumnFilter,
        filter: "includes",
      },
      {
        Header: "Summary",
        accessor: "summary",
        idAccessor: "issue_id",
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
      },
      {
        Header: "Due",
        accessor: "due_time",
      },
    ],
    []
  );
  return <Table columns={columns} data={[]} />;
};

export default ManageUser;
