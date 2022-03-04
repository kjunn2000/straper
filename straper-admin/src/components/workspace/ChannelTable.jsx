import React, { useMemo } from "react";
import { DateCell, IsDefaultCell } from "../../shared/table/TableCell";
import PaginationTable from "../../shared/table/PaginationTable";

const ChannelTable = ({ channelData }) => {
  const columns = useMemo(
    () => [
      {
        Header: "ID",
        accessor: "channel_id",
      },
      {
        Header: "Name",
        accessor: "channel_name",
      },
      {
        Header: "Creator ID",
        accessor: "creator_id",
      },
      {
        Header: "Created",
        accessor: "created_date",
        Cell: DateCell,
      },
      {
        Header: "Default",
        accessor: "is_default",
        Cell: IsDefaultCell,
      },
    ],
    []
  );

  const data = useMemo(() => channelData);

  return (
    <div>
      <PaginationTable columns={columns} data={data} />
    </div>
  );
};

export default ChannelTable;
