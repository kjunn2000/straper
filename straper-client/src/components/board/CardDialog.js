import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useState, useRef, useEffect } from "react";
import {
  MdOutlineTitle,
  MdOutlineDescription,
  MdLowPriority,
  MdRemoveCircle,
} from "react-icons/md";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import { useForm } from "react-hook-form";
import { BsCardChecklist } from "react-icons/bs";
import { AiFillDelete, AiOutlineClose } from "react-icons/ai";
import CardComment from "./CardComment";
import ActionDialog from "../dialog/ActionDialog";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import AddMember from "./AddMember";
import useWorkspaceStore from "../../store/workspaceStore";
import { iconStyle } from "../../utils/style/icon";

const CardDialog = ({ open, closeModal, card }) => {
  const board = useBoardStore((state) => state.board);
  const { register, handleSubmit, setValue } = useForm();
  let initialFocus = useRef(null);

  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [dueDate, setDueDate] = useState(new Date(card.due_date));
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

  const moreActionBtn = (text, action, Icon) => (
    <button
      className="w-full bg-indigo-400 text-white hover:bg-indigo-600 rounded shadow-lg shadow-indigo-500/50"
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

  return (
    <>
      <Transition appear show={open} as={Fragment}>
        <Dialog
          as="div"
          className="fixed inset-0 z-10 overflow-y-auto"
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
              <div className="inline-block p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl ">
                <div className="w-full flex justify-end hover:cursor-pointer">
                  <AiOutlineClose size={30} onClick={() => close()} />
                </div>
                <div
                  ref={initialFocus}
                  className="grid grid-cols-5 gap-x-8 gap-y-4"
                >
                  <form
                    onSubmit={handleSubmit(onSave)}
                    className="col-span-4 rounded-lg flex flex-col space-y-5 justify-center self-center"
                  >
                    <div className="flex space-x-2">
                      <MdOutlineTitle size={30} />
                      <input
                        className="p-1 rounded-lg hover:bg-gray-300 text-2xl font-bold"
                        defaultValue={card.title}
                        {...register("title")}
                      />
                    </div>
                    <div className="flex flex-col">
                      <div className="flex py-3 space-x-3">
                        <MdOutlineDescription size={30} />
                        <span className="font-semibold text-lg">
                          DESCRIPTION
                        </span>
                      </div>
                      <textarea
                        className="p-1 rounded-lg bg-gray-200 hover:bg-gray-300"
                        defaultValue={card.description}
                        {...register("description")}
                      />
                    </div>
                    <div className="grid grid-cols-5 gap-x-8 gap-y-4">
                      <div className="col-span-3 flex self-center py-3 space-x-3">
                        <MdLowPriority size={30} />
                        <span className="font-semibold text-lg">PRIORITY</span>
                      </div>
                      <select
                        defaultValue={card.priority}
                        {...register("priority")}
                        className="col-span-2 rounded-lg w-full hover:bg-gray-300 hover:cursor-pointer"
                      >
                        <option value="NO">No</option>
                        <option value="LOW">Low</option>
                        <option value="MEDIUM">Medium</option>
                        <option value="HIGH">High</option>
                      </select>
                    </div>
                    <button
                      type="submit"
                      className="self-end bg-indigo-400 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded-full"
                    >
                      SAVE
                    </button>
                  </form>

                  <div className="col-span-1">
                    <div>
                      <div className="flex py-3 space-x-3">
                        <span className="font-semibold text-gray-400">
                          DUE DATE
                        </span>
                      </div>
                      <DatePicker
                        selected={dueDate}
                        onChange={(date) => handleDueDateUpdate(date)}
                        className="p-1 rounded-lg hover:bg-gray-300"
                      />
                    </div>
                    <div>
                      <div className="flex py-3 space-x-3">
                        <span className="font-semibold text-gray-400">
                          MEMBER
                        </span>
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
                    <div className="flex flex-col space-y-2">
                      <div className="font-semibold text-gray-400 py-3">
                        MORE ACTION
                      </div>
                      <div className="flex flex-col space-y-5">
                        {moreActionBtn("CHECKLIST", () => {}, BsCardChecklist)}
                      </div>
                      <div className="flex flex-col space-y-5">
                        {moreActionBtn(
                          "DELETE CARD",
                          () => {
                            setIsDeleteDialogOpen(true);
                          },
                          AiFillDelete
                        )}
                      </div>
                    </div>
                  </div>
                </div>
                <CardComment />
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
