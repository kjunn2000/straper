import { ErrorMessage } from "@hookform/error-message";
import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import {
  useHistory,
  useParams,
} from "react-router-dom/cjs/react-router-dom.min";
import api from "../axios/api";
import SimpleDialog from "../shared/dialog/SimpleDialog";
import Tabs from "../components/workspace/Tabs";

const EditWorkspace = () => {
  const {
    handleSubmit,
    register,
    formState: { errors },
    reset,
  } = useForm();

  const { workspaceId } = useParams();
  const [workspace, setWorkspace] = useState();
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");
  const history = useHistory();

  useEffect(() => {
    fetchWorkspace();
  }, [workspaceId]);

  useEffect(() => resetForm(), [workspace]);

  const fetchWorkspace = async () => {
    const res = await api.get(`/protected/workspace/read/${workspaceId}`);
    if (res.data.Success) {
      setWorkspace(res.data.Data);
    }
  };

  const isFormDirty = (data) => {
    return !(data.workspace_name === workspace.workspace_name);
  };

  const onUpdate = (data) => {
    if (!isFormDirty(data)) {
      return;
    }
    api
      .post("/protected/workspace/update", data)
      .then((res) => {
        if (res.data.Success) {
          setWorkspace((prev) => ({
            ...prev,
            workspace_name: data.workspace_name,
          }));
          setShowSuccessDialog(true);
        } else {
          setDialogErrMsg("Something went wrong. Please try again.");
          setShowFailDialog(true);
        }
      })
      .catch((err) => {
        setShowFailDialog(true);
      });
  };

  const resetForm = () => {
    if (!workspace) {
      return;
    }
    reset({
      workspace_name: workspace.workspace_name,
    });
  };

  return (
    <div className="flex p-5">
      {workspace && (
        <>
          <form
            className="h-auto rounded-lg flex flex-col space-y-5 p-5 bg-white"
            onSubmit={handleSubmit(onUpdate)}
          >
            <div>
              <div className="text-xl text-center text-gray-500 font-semibold">
                Workspace Information
              </div>
            </div>
            <div>
              <div>Workspace ID</div>
              <input
                className="bg-gray-200 p-2 rounded-lg"
                defaultValue={workspace.workspace_id}
                readOnly
                {...register("workspace_id", {
                  required: "Id is required.",
                })}
              />
            </div>
            <div>
              <div>Workspace Name</div>
              <input
                className="bg-gray-100 p-2 rounded-lg"
                defaultValue={workspace.workspace_name}
                {...register("workspace_name", {
                  required: "Workspace name is required.",
                  minLength: {
                    value: 4,
                    message: "Workspace name should be at least 4 digits.",
                  },
                })}
              />
              <ErrorMessage
                errors={errors}
                name="workspace_name"
                as="p"
                className="text-red-500"
              />
              <span
                className="text-sm text-blue-600 hover:text-blue-300 hover:cursor-pointer text-underline self-end p-1"
                onClick={() => resetForm()}
              >
                RESET
              </span>
            </div>
            <div>
              <div>Creator ID</div>
              <input
                className="bg-gray-200 p-2 rounded-lg"
                defaultValue={workspace.creator_id}
                readOnly
              />
              <span
                className="text-sm text-blue-600 hover:text-blue-300 hover:cursor-pointer text-underline self-end p-1"
                onClick={() =>
                  history.push(`/manage/user/${workspace.creator_id}`)
                }
              >
                VIEW
              </span>
            </div>
            <button
              type="submit"
              className="bg-indigo-400 hover:bg-indigo-200 self-center w-48 p-1 rounded text-white"
            >
              CONFIRM UPDATE
            </button>
          </form>
          <div className="p-3 w-3/4">
            <Tabs
              userData={workspace.user_list}
              channelData={workspace.channel_list}
            />
          </div>
          <SimpleDialog
            isOpen={showSuccessDialog}
            setIsOpen={setShowSuccessDialog}
            title="Update Successfully"
            content="Workspace's lastest information is saved to database."
            buttonText="Close"
            buttonAction={() => setShowSuccessDialog(false)}
            buttonStatus="success"
          />

          <SimpleDialog
            isOpen={showFailDialog}
            setIsOpen={setShowFailDialog}
            title="Update Fail"
            content={dialogErrMsg}
            buttonText="Close"
            buttonStatus="fail"
          />
        </>
      )}
    </div>
  );
};

export default EditWorkspace;
