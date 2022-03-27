import { Dialog, Transition } from "@headlessui/react";
import { ErrorMessage } from "@hookform/error-message";
import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import { Fragment } from "react/cjs/react.production.min";
import api from "../../axios/api";
import useWorkspaceStore from "../../store/workspaceStore";

const EditWorkspaceDialog = ({ isOpen, close, workspace }) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();

  let editDialog = useRef(null);

  const updateWorkspace = useWorkspaceStore((state) => state.updateWorkspace);

  const resetForm = () => {
    if (!workspace) {
      return;
    }
    reset({
      workspace_name: workspace.workspace_name,
    });
  };

  const closeDialog = () => {
    resetForm();
    close();
  };

  const handleUpdate = (data) => {
    api.post("/protected/workspace/update", data).then((res) => {
      if (res.data.Success) {
        updateWorkspace(data);
        close();
      }
    });
  };

  return (
    <>
      {workspace && (
        <Transition appear show={isOpen} as={Fragment}>
          <Dialog
            as="div"
            className="fixed inset-0 z-20 overflow-y-auto"
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
                    onSubmit={handleSubmit(handleUpdate)}
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
                        className="bg-gray-100 p-2 rounded-lg w-full"
                        defaultValue={workspace.workspace_name}
                        {...register("workspace_name", {
                          required: "Workspace name is required.",
                          minLength: {
                            value: 4,
                            message:
                              "Workspace name should be at least 4 digits.",
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

export default EditWorkspaceDialog;
