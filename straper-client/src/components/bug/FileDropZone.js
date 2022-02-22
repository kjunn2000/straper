import React, { useCallback, useEffect, useMemo, useState } from "react";
import { useDropzone } from "react-dropzone";
import {
  base64ToArray,
  createBlobFile,
  downloadBlobFile,
  getAsByteArray,
} from "../../service/file";
import FileMessage from "../chat/FileMessage";
import api from "../../axios/api";
import useIssueStore from "../../store/issueStore";

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

function FileDropZone({ issueId, attachments }) {
  const [files, setFiles] = useState([]);

  const addIssueAttachments = useIssueStore(
    (state) => state.addIssueAttachments
  );

  useEffect(() => {
    if (!attachments) {
      return;
    }
    setFiles(
      attachments.map((attachment) => {
        const blob = createBlobFile(
          base64ToArray(attachment.file_bytes),
          attachment.file_type
        );
        return {
          ...attachment,
          blob,
        };
      })
    );
  }, [attachments]);

  const onDrop = useCallback(async (acceptedFiles) => {
    console.log(acceptedFiles);
    const attachments = acceptedFiles.map(async (file) => {
      const result = await getAsByteArray();
      return {
        file_name: file.name,
        file_type: file.type,
        file_bytes: Array.from(result),
      };
    });
    const payload = {
      issue_id: issueId,
      attachments,
    };
    const res = await api.post("/protected/issue/attachments/upload", payload);
    if (res.data.Success) {
      addIssueAttachments();
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

  const thumbs = files.map((file) => (
    <div onClick={() => downloadBlobFile(file.blob, file.file_name)}>
      <FileMessage file={file} />
    </div>
  ));

  return (
    <section>
      <div {...getRootProps({ style })}>
        <input {...getInputProps()} />
        <div>Drag and drop your files here.</div>
      </div>
      <aside>{thumbs}</aside>
    </section>
  );
}

export default FileDropZone;
