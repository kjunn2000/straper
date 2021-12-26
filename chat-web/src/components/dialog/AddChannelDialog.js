import { Dialog, Transition } from "@headlessui/react";
import { ErrorMessage } from "@hookform/error-message";
import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import { Fragment } from "react/cjs/react.production.min";
import api from "../../axios/api";
import useWorkspaceStore from "../../store/workspaceStore";

const AddChannelDialog = ({ isOpen, close }) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const setSelectedChannelIds = useWorkspaceStore((state) => state.setSelectedChannelIds);
  const addChannelToWorkspace = useWorkspaceStore((state) => state.addChannelToWorkspace);

  let addChannelBtn = useRef(null);

  const history = useHistory();

  const addNewChannel= async (data) => {
    const dto = {
      "workspace_id" : currWorkspace.workspace_id,
      "channel_name" : data?.channel_name
    }
    const res = await api.post("/protected/channel/create", dto);
    const channel = res.data.Data;
    addChannelToWorkspace(currWorkspace.workspace_id, channel);
    setCurrWorkspace(currWorkspace.workspace_id);
    setCurrChannel(channel.channel_id);
    setSelectedChannelIds(currWorkspace.workspace_id, channel.channel_id);
    closeDialog();
    history.push(`/channel/${currWorkspace.workspace_id}/${channel.channel_id}`)
  };

  const closeDialog = () => {
    reset();
    close();
  };

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog
        as="div"
        className="fixed inset-0 z-10 overflow-y-auto"
        onClose={closeDialog}
        initialFocus={addChannelBtn}
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
                Create Your Own Channel 
              </Dialog.Title>
              <form className="mt-2" onSubmit={handleSubmit(addNewChannel)}>
                <div className="self-center space-y-5">
                  <div>New Channel Name</div>
                  <input
                    className="bg-gray-200 p-2 w-full"
                    {...register("channel_name", {
                      required: {
                        value: true,
                        message: "Channel name cannot be empty.",
                      },
                    })}
                  />
                </div>
                <ErrorMessage errors={errors} name="channel_name" as="p" />

                <div className="mt-4 flex justify-end">
                  <button
                    type="submit"
                    className="inline-flex justify-center px-4 py-2 text-sm font-medium text-blue-900 bg-purple-300 border border-transparent rounded-md hover:bg-purple-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-purple-500"
                    ref={addChannelBtn}
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

export default AddChannelDialog;
