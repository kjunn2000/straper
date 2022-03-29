import React, { useMemo, useState } from "react";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import ActionDialog from "../../shared/dialog/ActionDialog";
import SimpleDialog from "../../shared/dialog/SimpleDialog";
import PaginationTable from "../../shared/table/PaginationTable";
import { ActionCell } from "../../shared/table/TableCell";

const UserTable = ({ userData, creatorId, handleRemoveUser }) => {
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [toDeleteUserId, setToDeleteUserId] = useState();
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");
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
        Header: "Action",
        idAccessor: "user_id",
        Cell: ActionCell,
        editAction: (userId) => {
          history.push(`/manage/user/${userId}`);
        },
        deleteAction: (userId) => {
          if (userId === creatorId) {
            setDialogErrMsg("Cannot remove the creator of workspace.");
            setShowFailDialog(true);
            return;
          }
          setToDeleteUserId(userId);
          setDeleteWarningDialogOpen(true);
        },
      },
    ],
    []
  );

  const data = useMemo(() => userData);

  return (
    <div>
      <PaginationTable columns={columns} data={data} />

      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Remove User Confirmation"
        content="Please confirm that the removed user will not able to be recovered."
        buttonText="Remove Anyway"
        buttonStatus="fail"
        buttonAction={() => handleRemoveUser(toDeleteUserId)}
        closeButtonText="Close"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Update Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default UserTable;
