import React, { useEffect, useMemo, useState } from "react";
import Pagination from "../shared/table/Pagination";
import Table, {
  ActionCell,
  DateCell,
  SelectColumnFilter,
  StatusPill,
} from "../shared/table/Table";
import api from "../axios/api";
import ActionDialog from "../shared/dialog/ActionDialog";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";

const ManageUser = () => {
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [toDeleteUserId, setToDeleteUserId] = useState();
  const [refreshPage, doRefreshPage] = useState(0);
  const history = useHistory();

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
        editAction: (userId) => {
          history.push(`/manage/user/${userId}`);
        },
        deleteAction: (userId) => {
          setToDeleteUserId(userId);
          setDeleteWarningDialogOpen(true);
        },
      },
    ],
    []
  );
  const [pageData, setPageData] = useState({
    isLoading: false,
    rowData: [],
    totalUsers: 0,
    searchStr: "",
  });

  useEffect(() => {
    fetchData(false);
  }, []);

  const fetchData = async (isNext, reload, searchStr) => {
    var cursor = "";
    if (!reload && pageData.rowData && pageData.rowData.length > 0) {
      if (isNext) {
        cursor = pageData.rowData[pageData.rowData.length - 1].user_id;
      } else {
        cursor = pageData.rowData[0].user_id;
      }
    }

    const newSearchStr = reload ? searchStr : pageData.searchStr;

    setPageData((prevState) => ({
      ...prevState,
      rowData: [],
      isLoading: true,
      totalUsers: 0,
      searchStr: newSearchStr,
    }));

    const res = await api.get(
      `/protected/user/list?limit=10&cursor=${cursor}&isNext=${isNext}&searchStr=${newSearchStr}`
    );
    if (res.data.Success) {
      const data = res.data.Data;
      if (!data.users && data.total_users === 0) {
        setPageData((prevState) => ({
          ...prevState,
          isLoading: false,
        }));
        return;
      }
      setPageData((prevState) => ({
        ...prevState,
        isLoading: false,
        rowData: data.users,
        totalUsers: data.total_users,
      }));
    }
  };

  const handleSearch = (value) => {
    fetchData(false, true, value);
    doRefreshPage((prev) => prev + 1);
  };

  const handleDeleteUser = async (userId) => {
    if (!userId || userId === "") {
      return;
    }
    const res = await api.post(`/protected/user/delete/${userId}`);
    if (res.data.Success) {
      const newData = pageData.rowData.filter(
        (user) => user.user_id !== userId
      );
      setPageData((prevState) => ({
        ...prevState,
        rowData: newData,
        totalUsers: prevState.totalUsers - 1,
      }));
    }
  };

  return (
    <div>
      <Table
        columns={columns}
        data={pageData.rowData}
        isLoading={pageData.isLoading}
        totalCount={pageData.totalUsers}
        onSearch={handleSearch}
      />
      <Pagination
        totalRows={pageData.totalUsers}
        pageChangeHandler={fetchData}
        rowsPerPage={10}
        refreshPage={refreshPage}
      />
      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Delete User Confirmation"
        content="Please confirm that the deleted user will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={() => handleDeleteUser(toDeleteUserId)}
        closeButtonText="Close"
      />
    </div>
  );
};

export default ManageUser;
