import React, { useState } from "react";
import { AiFillDelete } from "react-icons/ai";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import api from "../../axios/api";
import ActionDialog from "../../shared/dialog/ActionDialog";
import useIssueStore from "../../store/issueStore";

const DeleteIssueBtn = ({ issueId }) => {
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const deleteStateIssue = useIssueStore((state) => state.deleteIssue);
  const history = useHistory();

  const deleteIssue = async () => {
    const res = await api.post(`/protected/issue/delete/${issueId}`);
    if (res.data.Success) {
      deleteStateIssue(issueId);
      setDeleteWarningDialogOpen(false);
      history.push("/bug");
    }
  };

  return (
    <div>
      <button
        className="bg-red-500 text-white hover:bg-red-400 flex items-center rounded px-2 py-1"
        onClick={() => setDeleteWarningDialogOpen(true)}
      >
        <AiFillDelete size={20} />
        Delete
      </button>
      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Delete Issue Confirmation"
        content="Please confirm that the deleted issue will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={deleteIssue}
        closeButtonText="Close"
      />
    </div>
  );
};

export default DeleteIssueBtn;
