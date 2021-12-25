import { Dialog, Transition } from "@headlessui/react";
import { ErrorMessage } from "@hookform/error-message";
import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import { Fragment } from "react/cjs/react.production.min";
import api from "../../axios/api";
import useWorkspaceStore from "../../store/workspaceStore";

const AddWorkspaceDialog = ({ isOpen, close, toggleDialog }) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();
  const addWorkspace = useWorkspaceStore((state) => state.addWorkspace);
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const setSelectedChannelIds = useWorkspaceStore((state) => state.setSelectedChannelIds);

  let addWorkspaceBtn = useRef(null);

  const history = useHistory();

  const addNewWorkspace = (data) => {
    api.post("/protected/workspace/create", data).then((res) => {
      if (res.data.Success) {
        const newWorkspace = res.data.Data;
        addWorkspace(newWorkspace);
        setCurrWorkspace(newWorkspace.workspace_id)
        const channelId = newWorkspace.channel_list[0].channel_id
        setCurrChannel(channelId)
        setSelectedChannelIds(newWorkspace.workspace_id, channelId)
        history.push(`/channel/${newWorkspace.workspace_id}/${channelId}`)
      }
    });
    closeDialog();
  };

  const closeDialog = () => {
    reset();
    close();
  };

  const toggle = () => {
    reset();
    toggleDialog();
  };

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog
        as="div"
        className="fixed inset-0 z-10 overflow-y-auto"
        onClose={closeDialog}
        initialFocus={addWorkspaceBtn}
      >
        <div className="min-h-screen px-4 text-center">
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <Dialog.Overlay className="fixed inset-0" />
          </Transition.Child>

          <span
            className="inline-block h-screen align-middle"
            aria-hidden="true"
          >
            {" "}
            &#8203;{" "}
          </span>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <div className="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl space-y-5">
              <Dialog.Title
                as="h3"
                className="text-lg font-medium leading-6 text-gray-900"
              >
                Create Your Own Workspace
              </Dialog.Title>
              <form className="mt-2" onSubmit={handleSubmit(addNewWorkspace)}>
                <div className="self-center space-y-5">
                  <div>New Workspace Name</div>
                  <input
                    className="bg-gray-200 p-2 w-full"
                    {...register("workspace_name", {
                      required: {
                        value: true,
                        message: "Workspace name cannot be empty.",
                      },
                    })}
                  />
                </div>
                <ErrorMessage errors={errors} name="workspace_name" as="p" />
                <div
                  className="text-indigo-500 self-center cursor-pointer hover:text-indigo-300"
                  onClick={toggle}
                >
                  Join a workspace?
                </div>

                <div className="mt-4 flex justify-end">
                  <button
                    type="submit"
                    className="inline-flex justify-center px-4 py-2 text-sm font-medium text-blue-900 bg-purple-300 border border-transparent rounded-md hover:bg-purple-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-purple-500"
                    ref={addWorkspaceBtn}
                  >
                    Add
                  </button>
                </div>
              </form>
            </div>
          </Transition.Child>
        </div>
        <div className="absolute top-1 right-1">
          <button
            type="button"
            className="inline-flex justify-center px-2 py-1 text-sm font-medium text-gray-200 bg-gray-900 border border-transparent rounded hover:bg-gray-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-blue-500"
            onClick={closeDialog}
          >
            X
          </button>
        </div>
      </Dialog>
    </Transition>
  );
};

export default AddWorkspaceDialog;
