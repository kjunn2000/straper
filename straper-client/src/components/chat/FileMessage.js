import React from "react";
import { FcDocument } from "react-icons/fc";
import { AiOutlineDownload } from "react-icons/ai";
import { darkGrayBg } from "../../utils/style/color";

const FileMessage = ({ file }) => {
  return (
    <div
      className="flex shadow-2xl bg-gray-600 rounded px-3 py-1 justify-center items-center"
      style={darkGrayBg}
    >
      <FcDocument size={40} />
      <div className="px-5">
        <div>{file.file_name}</div>
        <div>{file.blob && file.blob.size / 1000} KB</div>
      </div>
      <AiOutlineDownload size={40} className="text-indigo-200" />
    </div>
  );
};

export default FileMessage;
