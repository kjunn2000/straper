import React, { useState } from "react";
import { useRef } from "react/cjs/react.development";
import { ImAttachment } from "react-icons/im";
import SimpleDialog from "../dialog/SimpleDialog";

const UploadButton = ({ handleFileAction }) => {
  const hiddenFileInput = useRef(null);
  const [isFailedDialogOpen, setFailedDialogOpen] = useState(false);

  const handleClick = (event) => {
    hiddenFileInput.current.click();
  };

  const handleChange = (event) => {
    const fileUploaded = event.target.files[0];
    const fileSize = fileUploaded.size;
    if (fileSize > 2000000) {
      setFailedDialogOpen(true);
      return;
    }
    handleFileAction(fileUploaded);
  };

  return (
    <>
      <button
        type="button"
        className="inline-flex items-center justify-center text-gray-500 px-5"
        onClick={handleClick}
      >
        <ImAttachment size="25" />
      </button>
      <input
        type="file"
        ref={hiddenFileInput}
        onChange={handleChange}
        style={{ display: "none" }}
      />
      <SimpleDialog
        isOpen={isFailedDialogOpen}
        setIsOpen={setFailedDialogOpen}
        title="File Size Exceed 2 MB"
        content="Failed to upload file. Please make sure the file size within 2 MB."
        buttonText="Close"
        buttonStatus="Fail"
      />
    </>
  );
};
export default UploadButton;
