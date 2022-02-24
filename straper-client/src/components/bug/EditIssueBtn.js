import React, { useState } from "react";
import { AiFillEdit } from "react-icons/ai";
import IssueDialog from "./IssueDialog";

const EditIssueBtn = ({ issue, setIssue }) => {
  const [editIssueDialogOpen, setEditIssueDialogOpen] = useState(false);

  return (
    <div>
      <button
        className="bg-gray-100 hover:bg-gray-200 flex items-center rounded px-2 py-1"
        onClick={() => setEditIssueDialogOpen(true)}
      >
        <AiFillEdit size={20} />
        Edit
      </button>
      <IssueDialog
        isOpen={editIssueDialogOpen}
        closeDialog={() => setEditIssueDialogOpen(false)}
        issue={issue}
        setIssue={setIssue}
      />
    </div>
  );
};

export default EditIssueBtn;
