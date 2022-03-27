import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useState, useRef } from "react";
import {
  MdOutlineTitle,
  MdOutlineDescription,
  MdLowPriority,
  MdRemoveCircle,
} from "react-icons/md";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import { useForm } from "react-hook-form";
import {
  BsCardChecklist,
  BsFillChatDotsFill,
  BsCalendarDate,
  BsPeople,
} from "react-icons/bs";
import { AiFillDelete, AiOutlineClose, AiFillBug } from "react-icons/ai";
import { FiMoreHorizontal } from "react-icons/fi";
import CardComment from "./CardComment";
import ActionDialog from "../../shared/dialog/ActionDialog";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import AddMember from "./AddMember";
import useWorkspaceStore from "../../store/workspaceStore";
import Checklist from "./Checklist";
import { Link } from "react-router-dom";
import EditIssueLink from "./EditIssueLink";

const CardDialog = ({ open, closeModal, card }) => {
  let initialFocus = useRef(null);
  const { register, handleSubmit, setValue } = useForm();

  const [dueDate, setDueDate] = useState(new Date(card.due_date));
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [isChecklistOpen, setIsChecklistOpen] = useState(false);

  const board = useBoardStore((state) => state.board);
  const currAccountList = useWorkspaceStore((state) => state.currAccountList);

  const close = () => {
    setValue("title", card.title);
    setValue("description", card.description);
    setValue("priority", card.priority);
    setDueDate(new Date(card.due_date));

    closeModal();
  };

  const onSave = (data) => {
    const payload = {
      ...data,
      list_id: card.list_id,
      card_id: card.card_id,
    };
    sendBoardMsg("BOARD_UPDATE_CARD", board.workspace_id, payload);
  };

  const handleDelete = () => {
    const payload = {
      list_id: card.list_id,
      card_id: card.card_id,
    };
    sendBoardMsg("BOARD_DELETE_CARD", board.workspace_id, payload);
  };

  const moreActionBtn = (text, action, Icon, bg) => (
    <button
      className={`w-48 md:w-full text-white rounded shadow-lg ${bg}`}
      onClick={() => action()}
    >
      <div className="flex space-x-3 p-2">
        <Icon size={20} />
        <span>{text}</span>
      </div>
    </button>
  );

  const handleDueDateUpdate = (date) => {
    date.setHours(0, 0, 0, 0);
    setDueDate(date);
    const payload = {
      list_id: card.list_id,
      card_id: card.card_id,
      due_date: date.toJSON(),
    };
    sendBoardMsg("BOARD_UPDATE_CARD_DUE_DATE", board.workspace_id, payload);
  };

  const removeUserFromCard = (userId) => {
    const payload = {
      list_id: card.list_id,
      card_id: card.card_id,
      member_id: userId,
    };
    sendBoardMsg("BOARD_CARD_REMOVE_MEMBER", board.workspace_id, payload);
  };

  const header = (title, Icon) => (
    <div className="flex py-3 space-x-3 items-center">
      <Icon size={20} className="text-gray-500" />
      <span className="font-semibold text-lg">{title}</span>
    </div>
  );

  return (
    <>
      <Transition appear show={open} as={Fragment}>
        <Dialog
          as="div"
          className="fixed inset-0 z-20 overflow-y-auto"
          onClose={() => {}}
          initialFocus={initialFocus}
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
              <div className="inline-block p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-gray-300 shadow-xl rounded-2xl ">
                <div className="w-full flex justify-end hover:cursor-pointer">
                  <AiOutlineClose
                    size={20}
                    onClick={() => close()}
                    className="text-gray-500"
                  />
                </div>
                <div
                  ref={initialFocus}
                  className="grid md:grid-cols-5 md:gap-x-8 md:gap-y-4"
                >
                  <div className="md:col-span-3 lg:col-span-4">
                    <form
                      onSubmit={handleSubmit(onSave)}
                      className="rounded-lg flex flex-col space-y-5 justify-center self-center"
                    >
                      <div className="flex space-x-2 items-center">
                        <MdOutlineTitle size={20} className="text-gray-500" />
                        <input
                          className="p-1 rounded-lg bg-gray-200 text-xl font-bold"
                          defaultValue={card.title}
                          {...register("title")}
                        />
                      </div>
                      <div className="flex flex-col">
                        {header("DESCRIPTION", MdOutlineDescription)}
                        <textarea
                          className="p-1 rounded-lg bg-gray-200"
                          defaultValue={card.description}
                          {...register("description")}
                        />
                      </div>
                      <div className="grid grid-cols-5 gap-x-8 gap-y-4">
                        {header("PRIORITY", MdLowPriority)}
                        <select
                          defaultValue={card.priority}
                          {...register("priority")}
                          className="col-span-2 rounded-lg w-full bg-gray-200 hover:cursor-pointer"
                        >
                          <option value="NO">No</option>
                          <option value="LOW">Low</option>
                          <option value="MEDIUM">Medium</option>
                          <option value="HIGH">High</option>
                        </select>
                      </div>
                      <button
                        type="submit"
                        className="self-end bg-indigo-600 hover:bg-indigo-400 text-white font-bold py-2 px-4 rounded-full"
                      >
                        SAVE
                      </button>
                    </form>

                    <Checklist
                      show={isChecklistOpen}
                      checklist={card.checklist}
                      listId={card.list_id}
                      cardId={card.card_id}
                    />
                  </div>
                  <div className="md:col-span-2 lg:col-span-1">
                    <div>
                      {header("DUE DATE", BsCalendarDate)}
                      <DatePicker
                        selected={dueDate}
                        onChange={(date) => handleDueDateUpdate(date)}
                        className="p-1 rounded-lg bg-gray-200"
                      />
                    </div>
                    <div>
                      <div className="flex items-center py-3 space-x-3">
                        {header("MEMBER", BsPeople)}
                        <AddMember card={card} />
                      </div>
                      <div>
                        <ul className="flex flex-col space-y-2">
                          {card.member_list &&
                            card.member_list.map((userId) => {
                              const user = currAccountList[userId];
                              return (
                                user && (
                                  <li
                                    key={user.user_id}
                                    className="group flex justify-between hover:bg-gray-200 
                                    rounded transition duration-300 p-2 font-semibold"
                                  >
                                    {user.username}
                                    <span
                                      className="opacity-0 group-hover:opacity-100 cursor-pointer pl-3"
                                      onClick={() => {
                                        removeUserFromCard(user.user_id);
                                      }}
                                    >
                                      <MdRemoveCircle
                                        size={25}
                                        className="text-red-500"
                                      />
                                    </span>
                                  </li>
                                )
                              );
                            })}
                        </ul>
                      </div>
                    </div>
                    <div>
                      <div className="flex items-center py-3 space-x-3">
                        {header("ISSUE LINK", AiFillBug)}
                        <EditIssueLink card={card} />
                      </div>
                      {card.issue_link && (
                        <div>
                          <Link
                            to={`/issue/${card.issue_link}`}
                            className="text-indigo-500 self-center cursor-pointer 
                          hover:text-indigo-300 hover:underline transition duration-150 font-semibold text-sm"
                          >
                            VIEW
                          </Link>
                        </div>
                      )}
                    </div>
                    <div className="flex flex-col space-y-2">
                      {header("MORE", FiMoreHorizontal)}
                      <div className="flex flex-col space-y-5">
                        {moreActionBtn(
                          isChecklistOpen ? "HIDE CHECKLIST" : "SHOW CHECKLIST",
                          () => {
                            setIsChecklistOpen((state) => !state);
                          },
                          BsCardChecklist,
                          "bg-indigo-600 hover:bg-indigo-400"
                        )}
                      </div>
                      <div className="flex flex-col space-y-5">
                        {moreActionBtn(
                          "DELETE CARD",
                          () => {
                            setIsDeleteDialogOpen(true);
                          },
                          AiFillDelete,
                          "bg-red-600 hover:bg-red-400"
                        )}
                      </div>
                    </div>
                  </div>
                </div>
                <div>
                  {header("ADD COMMENTS", BsFillChatDotsFill)}
                  <CardComment cardId={card.card_id} />
                </div>
              </div>
            </Transition.Child>
          </div>
        </Dialog>
      </Transition>
      <ActionDialog
        isOpen={isDeleteDialogOpen}
        setIsOpen={setIsDeleteDialogOpen}
        title="Confirm Delete Card"
        content="The card that you deleted cannot be recovered."
        buttonText="Delete"
        buttonStatus="fail"
        buttonAction={() => handleDelete()}
        closeButtonText="Close"
      />
    </>
  );
};

export default CardDialog;
