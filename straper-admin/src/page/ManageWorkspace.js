import React, { useEffect, useMemo, useState } from "react";
import Pagination from "../shared/table/Pagination";
import Table, { ActionCell, DateCell } from "../shared/table/Table";
import api from "../axios/api";
import ActionDialog from "../shared/dialog/ActionDialog";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";

const ManageWorkspace = () => {
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [toDeleteWorkspaceId, setToDeleteWorkspaceId] = useState();
  const [refreshPage, doRefreshPage] = useState(0);
  const history = useHistory();

  const columns = useMemo(
    () => [
      {
        Header: "ID",
        accessor: "workspace_id",
      },
      {
        Header: "Name",
        accessor: "workspace_name",
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
        Header: "Action",
        idAccessor: "workspace_id",
        Cell: ActionCell,
        editAction: (workspaceId) => {
          history.push(`/manage/workspace/${workspaceId}`);
        },
        deleteAction: (workspaceId) => {
          setToDeleteWorkspaceId(workspaceId);
          setDeleteWarningDialogOpen(true);
        },
      },
    ],
    []
  );
  const [pageData, setPageData] = useState({
    isLoading: false,
    rowData: [],
    totalWorkspaces: 0,
    searchStr: "",
  });

  useEffect(() => {
    fetchData(false);
  }, []);

  const fetchData = async (isNext, reload, searchStr) => {
    var cursor = "";
    if (!reload && pageData.rowData && pageData.rowData.length > 0) {
      if (isNext) {
        cursor = pageData.rowData[pageData.rowData.length - 1].cursor;
      } else {
        cursor = pageData.rowData[0].cursor;
      }
    }

    const newSearchStr = reload ? searchStr : pageData.searchStr;

    setPageData((prevState) => ({
      ...prevState,
      rowData: [],
      isLoading: true,
      totalWorkspaces: 0,
      searchStr: newSearchStr,
    }));

    const res = await api.get(
      `/protected/workspace/list?limit=10&cursor=${cursor}&isNext=${isNext}&searchStr=${newSearchStr}`
    );
    if (res.data.Success) {
      const data = res.data.Data;
      if (!data.workspaces && data.total_workspaces === 0) {
        setPageData((prevState) => ({
          ...prevState,
          isLoading: false,
        }));
        return;
      }
      setPageData((prevState) => ({
        ...prevState,
        isLoading: false,
        rowData: data.workspaces,
        totalWorkspaces: data.total_workspaces,
      }));
    }
  };

  const handleSearch = (value) => {
    fetchData(false, true, value);
    doRefreshPage((prev) => prev + 1);
  };

  const handleDeleteWorkspace = async (workspaceId) => {
    if (!workspaceId || workspaceId === "") {
      return;
    }
    const res = await api.post(`/protected/workspace/delete/${workspaceId}`);
    if (res.data.Success) {
      const newData = pageData.rowData.filter(
        (workspace) => workspace.workspace_id !== workspaceId
      );
      setPageData((prevState) => ({
        ...prevState,
        rowData: newData,
        totalWorkspaces: prevState.totalWorkspaces - 1,
      }));
    }
  };

  return (
    <div>
      <Table
        columns={columns}
        data={pageData.rowData}
        isLoading={pageData.isLoading}
        totalCount={pageData.totalWorkspaces}
        onSearch={handleSearch}
      />
      <Pagination
        totalRows={pageData.totalWorkspaces}
        pageChangeHandler={fetchData}
        rowsPerPage={10}
        refreshPage={refreshPage}
      />
      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Delete Workspace Confirmation"
        content="Please confirm that the deleted workspace will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={() => handleDeleteWorkspace(toDeleteWorkspaceId)}
        closeButtonText="Close"
      />
    </div>
  );
};

export default ManageWorkspace;
