import React from "react";
import { AiFillEdit } from "react-icons/ai";

const EditIssueBtn = () => {
  return (
    <button className="bg-gray-100 hover:bg-gray-200 flex items-center rounded px-2 py-1">
      <AiFillEdit size={20} />
      Edit
    </button>
  );
};

export default EditIssueBtn;
