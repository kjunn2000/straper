import React, { useCallback, useMemo, useState } from "react";
import { useDropzone } from "react-dropzone";
import { createBlobFile, downloadBlobFile } from "../../service/file";
import api from "../../axios/api";
import useIssueStore from "../../store/issueStore";
import { AiFillDelete } from "react-icons/ai";
import { BsDownload } from "react-icons/bs";
import { iconStyle } from "../../utils/style/icon";
import SimpleDialog from "../../shared/dialog/SimpleDialog";

const baseStyle = {
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  padding: "20px",
  borderWidth: 2,
  borderRadius: 2,
  borderColor: "#eeeeee",
  borderStyle: "dashed",
  backgroundColor: "#fafafa",
  color: "#bdbdbd",
  transition: "border .3s ease-in-out",
};

const activeStyle = {
  borderColor: "#2196f3",
};

const acceptStyle = {
  borderColor: "#00e676",
};

const rejectStyle = {
  borderColor: "#ff1744",
};

function FileDropZone({ issueId, attachments, getIssueData }) {
  const [successUploadOpen, setSuccessUploadOpen] = useState(false);
  const [successDeleteOpen, setSuccessDeleteOpen] = useState(false);

  const addIssueAttachments = useIssueStore(
    (state) => state.addIssueAttachments
  );

  const deleteAttachment = useIssueStore((state) => state.deleteAttachment);

  const downloadFile = async (file) => {
    const res = await api.get(
      `/protected/issue/attachment/download/${file.fid}`,
      { responseType: "blob" }
    );
    const blob = createBlobFile(res.data, file.file_type);
    downloadBlobFile(blob, file.file_name);
  };

  const deleteFile = async (fid) => {
    const res = await api.post(`/protected/issue/attachment/delete/${fid}`);
    if (res.data.Success) {
      deleteAttachment(issueId, fid);
      getIssueData();
      setSuccessDeleteOpen(true);
    }
  };

  const onDrop = useCallback(async (acceptedFiles) => {
    const formData = new FormData();
    acceptedFiles.forEach(async (file) => {
      formData.append("files", file, file.name);
      formData.append("types", file.type);
    });
    formData.append("issue_id", issueId);

    const res = await api.post(
      "/protected/issue/attachments/upload",
      formData,
      {
        headers: { "Content-Type": "multipart/form-data" },
      }
    );
    if (res.data.Success) {
      addIssueAttachments(issueId, res.data.Data);
      getIssueData();
      setSuccessUploadOpen(true);
    }
  }, []);

  const {
    getRootProps,
    getInputProps,
    isDragActive,
    isDragAccept,
    isDragReject,
  } = useDropzone({
    onDrop,
  });

  const style = useMemo(
    () => ({
      ...baseStyle,
      ...(isDragActive ? activeStyle : {}),
      ...(isDragAccept ? acceptStyle : {}),
      ...(isDragReject ? rejectStyle : {}),
    }),
    [isDragActive, isDragReject, isDragAccept]
  );

  const thumbs =
    attachments &&
    attachments.map((file) => (
      <div
        className="group bg-gray-100 rounded p-2 text-semibold text-sm italic flex justify-between"
        key={file.fid}
      >
        <span>{file.file_name}</span>
        <div className="flex">
          <span
            className="opacity-0 group-hover:opacity-100 cursor-pointer transition duration-150"
            onClick={() => downloadFile(file)}
          >
            <BsDownload style={iconStyle} className="text-indigo-800" />
          </span>
          <span
            className="opacity-0 group-hover:opacity-100 cursor-pointer pl-3 transition duration-150"
            onClick={() => deleteFile(file.fid)}
          >
            <AiFillDelete style={iconStyle} className="text-indigo-800" />
          </span>
        </div>
      </div>
    ));

  return (
    <section>
      <div {...getRootProps({ style })}>
        <input {...getInputProps()} />
        <div>Drag and drop your files here.</div>
      </div>
      <aside className="flex flex-col space-y-1 py-2">
        {attachments && thumbs}
      </aside>
      <SimpleDialog
        isOpen={successUploadOpen}
        setIsOpen={setSuccessUploadOpen}
        title="Success Upload Attachment"
        content="Successfully uploaded attachment to the issue."
        buttonText="Close"
        buttonStatus="success"
      />
      <SimpleDialog
        isOpen={successDeleteOpen}
        setIsOpen={setSuccessDeleteOpen}
        title="Success Delete Attachment"
        content="Successfully deleted attachment from the issue."
        buttonText="Close"
        buttonStatus="success"
      />
    </section>
  );
}

export default FileDropZone;
