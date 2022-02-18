import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";
import DatePicker from "react-datepicker";

export default function CreateIssueDialog({ isOpen, closeDialog }) {
  const cancelButtonRef = useRef();
  const [errMsg, setErrMsg] = useState("");
  const [dueDate, setDueDate] = useState();

  const {
    register,
    handleSubmit,
    resetField,
    clearErrors,
    formState: { errors },
  } = useForm();

  const onClose = () => {
    resetField();
    clearErrors();
    closeDialog();
  };

  const Title = ({ text, required }) => (
    <div className="font-semibold text-sm text-gray-500">
      {text}
      {required && <span className="text-red-500">*</span>}
    </div>
  );

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog
        as="div"
        className="fixed inset-0 z-10 overflow-y-auto"
        onClose={() => onClose()}
        initialFocus={cancelButtonRef}
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
            &#8203;
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
            <div className="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl">
              <Dialog.Title
                as="h3"
                className="text-lg font-medium leading-6 text-gray-900"
              >
                Create Issue
              </Dialog.Title>
              <form
                onSubmit={handleSubmit(() => {})}
                className="rounded-lg flex-col space-y-5 w-96 h-auto self-center py-5"
              >
                <div>
                  <Title text="Issue Type" required={true} />
                  <select
                    {...register("type", {
                      required: "Issue type is required.",
                    })}
                    className="w-1/2 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                  >
                    <option value="task">Task</option>
                    <option value="subtask">Subtask</option>
                    <option value="bug">Bug</option>
                    <option value="story">Story</option>
                    <option value="epic">Epic</option>
                  </select>
                  <ErrorMessage errors={errors} name="type" as="p" />
                </div>
                <div>
                  <Title text="Summary" required={true} />
                  <input
                    className="w-full bg-gray-200 p-2 rounded focus:outline-none"
                    {...register("summary", {
                      required: "Summary is required.",
                    })}
                  />
                  <ErrorMessage errors={errors} name="summary" as="p" />
                </div>
                <div>
                  <Title text="Description" required={false} />
                  <textarea
                    className="w-full p-1 rounded-lg bg-gray-200"
                    {...register("description")}
                  />
                  <ErrorMessage errors={errors} name="description" as="p" />
                </div>
                <div>
                  <Title text="Acceptance Criteria" required={false} />
                  <textarea
                    className="w-full p-1 rounded-lg bg-gray-200"
                    {...register("acceptance_criteria")}
                  />
                  <ErrorMessage
                    errors={errors}
                    name="acceptance_criteria"
                    as="p"
                  />
                </div>
                <div>
                  <Title text="Epic Link" required={false} />
                  <select
                    {...register("epic_link")}
                    className="w-2/3 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                  ></select>
                  <ErrorMessage errors={errors} name="epic_link" as="p" />
                </div>
                <div>
                  <Title text="Story Point" required={false} />
                  <input
                    type="number"
                    className="w-1/3 bg-gray-200 p-2 rounded focus:outline-none"
                    {...register("story_point")}
                  />
                  <ErrorMessage errors={errors} name="story_point" as="p" />
                </div>
                <div>
                  <Title text="Replicate Step" required={false} />
                  <textarea
                    className="w-full p-1 rounded-lg bg-gray-200"
                    {...register("replicate_step")}
                  />
                  <ErrorMessage errors={errors} name="replicate_step" as="p" />
                </div>
                <div>
                  <Title text="Environment" required={false} />
                  <input
                    className="w-full bg-gray-200 p-2 rounded focus:outline-none"
                    {...register("environment")}
                  />
                  <ErrorMessage errors={errors} name="environment" as="p" />
                </div>
                <div>
                  <Title text="Serverity" required={false} />
                  <select
                    {...register("serverity")}
                    className="w-2/3 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                  ></select>
                  <ErrorMessage errors={errors} name="serverity" as="p" />
                </div>
                <div>
                  <Title text="Label" required={false} />
                  <input
                    className="w-full bg-gray-200 p-2 rounded focus:outline-none"
                    {...register("label")}
                  />
                  <ErrorMessage errors={errors} name="label" as="p" />
                </div>
                <div>
                  <Title text="Assignee" required={false} />
                  <select
                    {...register("assignee")}
                    className="w-2/3 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                  ></select>
                  <ErrorMessage errors={errors} name="assignee" as="p" />
                </div>
                <div>
                  <Title text="Due Time" required={false} />
                  <DatePicker
                    selected={dueDate}
                    onChange={(date) => setDueDate(date)}
                    className="p-1 rounded-lg bg-gray-200"
                  />
                </div>

                {errMsg !== "" && (
                  <div className="text-red-600 self-center">{errMsg}</div>
                )}
                <div className="flex justify-end">
                  <button
                    type="submit"
                    className="bg-indigo-600 text-white self-center p-2 rounded-l
                    hover:bg-indigo-800"
                  >
                    Create
                  </button>
                  <button
                    type="button"
                    className="text-indigo-600 self-center p-2"
                    onClick={() => onClose()}
                  >
                    Cancel
                  </button>
                </div>
              </form>
            </div>
          </Transition.Child>
        </div>
      </Dialog>
    </Transition>
  );
}
