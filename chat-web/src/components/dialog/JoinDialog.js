import { Dialog, Transition } from "@headlessui/react";
import { ErrorMessage } from "@hookform/error-message";
import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import { Fragment } from "react/cjs/react.production.min";

const JoinDialog = ({ isOpen, close, toggleDialog, joinAction, type}) => {
  const {
    register,
    handleSubmit,
    reset,
    setError,
    formState: { errors },
  } = useForm();

  const joinBtn = useRef(null);

  const closeDialog = () => {
    reset();
    close();
  };

  const toggle = () => {
    reset();
    toggleDialog();
  };

  const executeJoinActoin = async (data) => {
    const errMsg = await joinAction(data);
    if (!errMsg || errMsg == ""){
      closeDialog();
      return;
    }
    setError(type=="workspace" ? "workspace_id" : "channel_id",{
      type: "bad_request",
      message: errMsg,
    });
  };

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog
        as="div"
        className="fixed inset-0 z-10 overflow-y-auto"
        onClose={closeDialog}
        initialFocus={joinBtn}
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
                {type == "workspace" ?  "Join a Workspace" : "Join a Channel"}
              </Dialog.Title>
              <form className="mt-2" onSubmit={handleSubmit(executeJoinActoin)}>
                <div className="self-center space-y-5">
                  <div>
                    {
                      type == "workspace" ? "Workspace ID" : "Channel ID"
                    }
                  </div>
                  {
                    type == "workspace" ? 
                      <input
                        className="bg-gray-200 p-2 w-full"
                        {...register("workspace_id", {
                          required: {
                            value: true,
                            message: "Workspace ID cannot be empty.",
                          },
                        })}
                      /> :
                      <input
                        className="bg-gray-200 p-2 w-full"
                        {...register("channel_id", {
                          required: {
                            value: true,
                            message: "Channel ID cannot be empty.",
                          },
                        })}
                      />
                  }
                </div>
                {type=="workspace" ? 
                  <ErrorMessage errors={errors} name="workspace_id" as="p" />
                  :
                  <ErrorMessage errors={errors} name="channel_id" as="p" />
                }
                <div
                  className="text-indigo-500 self-center cursor-pointer hover:text-indigo-300"
                  onClick={toggle}
                >
                  {type == "workspace" ? "Create new workspace?" : "Create new channel?"}
                </div>

                <div className="mt-4 flex justify-end">
                  <button
                    type="submit"
                    className="inline-flex justify-center px-4 py-2 text-sm font-medium text-blue-900 bg-purple-300 border border-transparent rounded-md hover:bg-purple-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-purple-500"
                    ref={joinBtn}
                  >
                    Join
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

export default JoinDialog;
