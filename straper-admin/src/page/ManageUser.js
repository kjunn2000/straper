import React, { useEffect, useMemo, useState } from "react";
import Pagination from "../shared/table/Pagination";
import Table, {
  ActionCell,
  DateCell,
  SelectColumnFilter,
  StatusPill,
} from "../shared/table/Table";
import api from "../axios/api";

const ManageUser = () => {
  const columns = useMemo(
    () => [
      {
        Header: "ID",
        accessor: "user_id",
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
      {
        Header: "Status",
        accessor: "status",
        Cell: StatusPill,
      },
      {
        Header: "Created",
        accessor: "created_date",
        Cell: DateCell,
      },
      {
        Header: "Updated",
        accessor: "updated_date",
        Cell: DateCell,
      },
      {
        Header: "Action",
        idAccessor: "user_id",
        Cell: ActionCell,
      },
    ],
    []
  );
  const [pageData, setPageData] = useState({
    isLoading: false,
    rowData: [],
    totalUsers: 0,
  });

  useEffect(() => {
    fetchData("", false);
  }, []);

  const fetchData = async (cursor, isNext) => {
    setPageData((prevState) => ({
      ...prevState,
      rowData: [],
      isLoading: true,
      totalUsers: 0,
    }));

    const res = await api.get(
      `/protected/users?limit=10&cursor=${cursor}&isNext=${isNext}`
    );
    if (res.data.Success) {
      const data = res.data.Data;
      console.log(data.users);
      if (!data.users && data.total_users === 0) {
        setPageData((prevState) => ({
          ...prevState,
          isLoading: false,
        }));
        return;
      }
      setPageData({
        isLoading: false,
        rowData: data.users,
        totalUsers: data.total_users,
      });
    }
  };

  return (
    <div>
      <Table
        columns={columns}
        data={pageData.rowData}
        isLoading={pageData.isLoading}
      />
      <Pagination
        totalRows={pageData.totalUsers}
        pageChangeHandler={fetchData}
        rowsPerPage={10}
      />
    </div>
  );
};

export default ManageUser;
