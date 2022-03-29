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

  const handleRemoveUser = async (userId) => {
    if (!userId || userId === "") {
      return;
    }
    const res = await api.post(
      `/protected/workspace/remove/${workspace.workspace_id}/${userId}`
    );
    if (res.data.Success) {
      const newUserList = workspace.user_list.filter(
        (user) => user.user_id !== userId
      );
      console.log(newUserList);
      console.log(workspace);
      setWorkspace((prevState) => ({
        ...prevState,
        user_list: newUserList,
      }));
    }
  };

  const handleUpdateChannel = async (data) => {
    const res = await api.post("/protected/workspace/channel/update", data);
    if (res.data.Success) {
      const newChannelList = workspace.channel_list.map((channel) => {
        if (channel.channel_id !== data.channel_id) {
          return channel;
        }
        channel.channel_name = data.channel_name;
        return channel;
      });
      setWorkspace((prevState) => ({
        ...prevState,
        channel_list: newChannelList,
      }));
    }
  };

  const handleDeleteChannel = async (channelId) => {
    const res = await api.post(
      `/protected/workspace/channel/delete/${channelId}`
    );
    if (res.data.Success) {
      const newChannelList = workspace.channel_list.filter(
        (channel) => channel.channel_id !== channelId
      );
      setWorkspace((prevState) => ({
        ...prevState,
        channel_list: newChannelList,
      }));
    }
  };

  return (
    <>
      {workspace && (
        <div className="flex flex-col lg:flex-row p-5">
          <form
            className="w-2/3 lg:w-1/3 h-auto rounded-lg flex flex-col space-y-5 p-5 bg-white"
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
                className="bg-gray-200 p-2 rounded-lg w-full"
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
                className="w-full bg-gray-100 p-2 rounded-lg"
                defaultValue={workspace.workspace_name}
                {...register("workspace_name", {
                  required: "Workspace name is required.",
                  minLength: {
                    value: 4,
                    message: "Workspace name at least 4 chars.",
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
                className="w-full bg-gray-200 p-2 rounded-lg"
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
          <div className="lg:w-2/3 pt-8 pb-16 md:pb-8 lg:p-3">
            <Tabs
              userData={workspace.user_list}
              channelData={workspace.channel_list}
              creatorId={workspace.creator_id}
              handleRemoveUser={handleRemoveUser}
              handleUpdateChannel={handleUpdateChannel}
              handleDeleteChannel={handleDeleteChannel}
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
        </div>
      )}
    </>
  );
};

export default EditWorkspace;
