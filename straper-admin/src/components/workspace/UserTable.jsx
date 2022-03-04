import React, { useMemo } from "react";
import PaginationTable from "../../shared/table/PaginationTable";
import { DateCell, StatusPill, UserIdCell } from "../../shared/table/TableCell";

const UserTable = ({ userData }) => {
  const columns = useMemo(
    () => [
      {
        Header: "ID",
        accessor: "user_id",
        Cell: UserIdCell,
      },
      {
        Header: "Username",
        accessor: "username",
      },
      {
        Header: "Email",
        accessor: "email",
      },
      {
        Header: "Phone No",
        accessor: "phone_no",
      },
    ],
    []
  );

  const data = useMemo(() => userData);

  return (
    <div>
      <PaginationTable columns={columns} data={data} />
    </div>
  );
};

export default UserTable;
