import { Dialog, Transition } from "@headlessui/react";
import { ErrorMessage } from "@hookform/error-message";
import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import { Fragment } from "react/cjs/react.production.min";

const EditChannelDialog = ({ isOpen, close, handleUpdateChannel, channel }) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  let editDialog = useRef(null);

  const resetForm = () => {
    if (!channel) {
      return;
    }
    reset({
      channel_name: channel.channel_name,
    });
  };

  const closeDialog = () => {
    resetForm();
    close();
  };

  const updateChannel = (data) => {
    handleUpdateChannel(data);
    close();
  };

  return (
    <>
      {" "}
      {channel && (
        <Transition appear show={isOpen} as={Fragment}>
          <Dialog
            as="div"
            className="fixed inset-0 z-10 overflow-y-auto"
            onClose={closeDialog}
            initialFocus={editDialog}
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
                  <form
                    className="h-auto rounded-lg flex flex-col space-y-5 p-5 bg-white"
                    onSubmit={handleSubmit(updateChannel)}
                  >
                    <div>
                      <div className="text-xl text-center text-gray-500 font-semibold">
                        Channel Information
                      </div>
                    </div>
                    <div>
                      <div>Channel ID</div>
                      <input
                        className="bg-gray-200 p-2 rounded-lg w-full"
                        defaultValue={channel.channel_id}
                        readOnly
                        {...register("channel_id", {
                          required: "Id is required.",
                        })}
                      />
                    </div>
                    <div>
                      <div>Channel Name</div>
                      <input
                        className="bg-gray-100 p-2 rounded-lg w-full"
                        defaultValue={channel.channel_name}
                        {...register("channel_name", {
                          required: "Channel name is required.",
                          minLength: {
                            value: 4,
                            message: "Channel name at least 4 chars.",
                          },
                        })}
                      />
                      <ErrorMessage
                        errors={errors}
                        name="channel_name"
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
                    <button
                      type="submit"
                      className="bg-indigo-400 hover:bg-indigo-200 self-center w-48 p-1 rounded text-white"
                    >
                      CONFIRM UPDATE
                    </button>
                  </form>
                </div>
              </Transition.Child>
            </div>
          </Dialog>
        </Transition>
      )}
    </>
  );
};

export default EditChannelDialog;
