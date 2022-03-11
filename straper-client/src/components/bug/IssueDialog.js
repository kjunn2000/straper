import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";
import DatePicker from "react-datepicker";
import useWorkpaceStore from "../../store/workspaceStore";
import api from "../../axios/api";
import useIdentityStore from "../../store/identityStore";
import useIssueStore from "../../store/issueStore";
import { removeEmptyFields } from "../../service/object";

export default function IssueDialog({ isOpen, closeDialog, issue, setIssue }) {
  const cancelButtonRef = useRef();
  const [errMsg, setErrMsg] = useState("");
  const [dueDate, setDueDate] = useState(new Date());
  const [epicLinkOptions, setEpicLinkOptions] = useState([]);
  const [assigneeOptions, setAssigneeOptions] = useState([]);
  const [editMode, setEditMode] = useState(issue != null);

  const currWorkspace = useWorkpaceStore((state) => state.currWorkspace);
  const addIssue = useIssueStore((state) => state.addIssue);
  const updateIssue = useIssueStore((state) => state.updateIssue);
  const setStateAssigneeOptions = useIssueStore(
    (state) => state.setAssigneeOptions
  );

  const {
    register,
    handleSubmit,
    reset,
    clearErrors,
    setValue,
    formState: { errors },
  } = useForm();

  useEffect(() => {
    fetchEpicLinkOptions();
    fetchAssigneeOptions();
  }, []);

  useEffect(() => {
    initFields();
  }, [issue]);

  const initFields = () => {
    if (!issue) {
      return;
    }
    setValue("type", issue.type);
    setValue("summary", issue.summary);
    setValue("description", issue.description);
    setValue("acceptance_criteria", issue.acceptance_criteria);
    setValue("epic_link", issue.epic_link);
    setValue("story_point", issue.story_point);
    setValue("replicate_step", issue.replicate_step);
    setValue("environment", issue.environment);
    setValue("workaround", issue.workaround);
    setValue("priority", issue.priority);
    setValue("serverity", issue.serverity);
    setValue("label", issue.label);
    setValue("assignee", issue.assignee);
    setDueDate(new Date(issue.due_time));
    setValue("status", issue.status);
  };

  const fetchEpicLinkOptions = async () => {
    const res = await api.get(
      `/protected/issue/epic-link/option/${currWorkspace.workspace_id}`
    );
    if (res.data.Success) {
      setEpicLinkOptions(res.data.Data);
    }
  };

  const fetchAssigneeOptions = async () => {
    const res = await api.get(
      `/protected/issue/assignee/option/${currWorkspace.workspace_id}`
    );
    if (res.data.Success) {
      setAssigneeOptions(res.data.Data);
      setStateAssigneeOptions(res.data.Data);
    }
  };

  const onClose = () => {
    if (!editMode) {
      reset();
      setDueDate(new Date());
      clearErrors();
    }
    closeDialog();
  };

  const createIssue = async (data) => {
    if (dueDate) {
      data.due_time = dueDate.toJSON();
    }
    data.workspace_id = currWorkspace.workspace_id;
    data.story_point = parseInt(data.story_point);
    removeEmptyFields(data);
    const res = await api.post("/protected/issue/create", data);
    if (res.data.Success) {
      addIssue(res.data.Data);
      onClose();
    }
  };

  const editIssue = async (data) => {
    if (dueDate) {
      data.due_time = dueDate.toJSON();
    }
    data.workspace_id = currWorkspace.workspace_id;
    data.story_point = parseInt(data.story_point);
    data.issue_id = issue.issue_id;
    removeEmptyFields(data);
    const res = await api.post("/protected/issue/update", data);
    if (res.data.Success) {
      updateIssue(res.data.Data);
      setIssue(res.data.Data);
      onClose();
    }
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
                {editMode ? "Edit Issue" : "Create Issue"}
              </Dialog.Title>
              <form
                onSubmit={handleSubmit((data) => {
                  editMode ? editIssue(data) : createIssue(data);
                })}
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
                  >
                    {epicLinkOptions &&
                      epicLinkOptions.map((option) => (
                        <option key={option.issue_id} value={option.issue_id}>
                          {option.summary}
                        </option>
                      ))}
                    <option value={""}>No</option>
                  </select>
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
                  <Title text="Workaround" required={false} />
                  <textarea
                    className="w-full p-1 rounded-lg bg-gray-200"
                    {...register("workaround")}
                  />
                  <ErrorMessage errors={errors} name="workaround" as="p" />
                </div>
                <div>
                  <Title text="Priority" required={false} />
                  <select
                    {...register("backlog_priority")}
                    className="w-2/3 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                  >
                    {Array.from({ length: 5 }, (v, k) => k + 1).map((val) => (
                      <option key={val} value={val}>
                        {val}
                      </option>
                    ))}
                  </select>
                  <ErrorMessage
                    errors={errors}
                    name="backlog_priority"
                    as="p"
                  />
                </div>
                <div>
                  <Title text="Serverity" required={false} />
                  <select
                    {...register("serverity")}
                    className="w-2/3 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                  >
                    <option value="blocker">Blocker</option>
                    <option value="critical">Critical</option>
                    <option value="major">Major</option>
                    <option value="minor">Minor</option>
                    <option value="low">Low</option>
                  </select>
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
                  >
                    {assigneeOptions &&
                      assigneeOptions.map((option) => (
                        <option key={option.user_id} value={option.user_id}>
                          {option.username}
                        </option>
                      ))}
                  </select>
                  <ErrorMessage errors={errors} name="assignee" as="p" />
                </div>
                <div>
                  <Title text="Due Time" required={false} />
                  <DatePicker
                    selected={dueDate}
                    onChange={(date) => setDueDate(date)}
                    className="p-1 rounded-lg bg-gray-200"
                    showTimeSelect
                    timeIntervals={15}
                    dateFormat="Pp"
                  />
                </div>
                {editMode && (
                  <div>
                    <Title text="Status" required={true} />
                    <select
                      {...register("status")}
                      className="w-2/3 p-2 rounded bg-gray-200 hover:cursor-pointer focus:outline-none"
                    >
                      <option value="ACTIVE">ACTIVE</option>
                      <option value="INACTIVE">INACTIVE</option>
                      <option value="CLOSED">CLOSED</option>
                    </select>
                    <ErrorMessage errors={errors} name="status" as="p" />
                  </div>
                )}

                {errMsg !== "" && (
                  <div className="text-red-600 self-center">{errMsg}</div>
                )}
                <div className="flex justify-end">
                  <button
                    type="submit"
                    className="bg-indigo-600 text-white self-center p-2 rounded-l
                    hover:bg-indigo-800"
                  >
                    {editMode ? "Edit" : "Create"}
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
