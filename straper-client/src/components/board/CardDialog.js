import { Dialog, Transition } from "@headlessui/react";
import { Fragment } from "react";
import {
  MdOutlineTitle,
  MdOutlineDescription,
  MdLowPriority,
} from "react-icons/md";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";
import { useForm } from "react-hook-form";
import { BsFillCalendarDateFill } from "react-icons/bs";
import { AiFillDelete } from "react-icons/ai";
import CardComment from "./CardComment";

const CardDialog = ({ open, closeModal, card }) => {
  const board = useBoardStore((state) => state.board);
  const { register, handleSubmit, setValue } = useForm();

  const close = () => {
    setValue("title", card.title);
    setValue("description", card.description);
    setValue("priority", card.priority);

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

  return (
    <>
      <Transition appear show={open} as={Fragment}>
        <Dialog
          as="div"
          className="fixed inset-0 z-10 overflow-y-auto"
          onClose={() => close()}
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
                <div className="flex space-x-5">
                  <form
                    onSubmit={handleSubmit(onSave)}
                    className="rounded-lg flex flex-col space-y-5 justify-center self-center"
                  >
                    <div className="flex flex-col">
                      <div className="flex p-3 space-x-3">
                        <MdOutlineTitle size={20} />
                        <span className="font-semibold text-sm">TITLE</span>
                      </div>
                      <input
                        className="p-1 rounded-lg hover:bg-gray-300"
                        defaultValue={card.title}
                        {...register("title")}
                      />
                    </div>
                    <div className="flex flex-col">
                      <div className="flex p-3 space-x-3">
                        <MdOutlineDescription size={20} />
                        <span className="font-semibold text-sm">
                          DESCRIPTION
                        </span>
                      </div>
                      <textarea
                        className="p-1 rounded-lg bg-gray-200 hover:bg-gray-300"
                        defaultValue={card.description}
                        {...register("description")}
                      />
                    </div>
                    <div className="flex">
                      <div className="flex self-center p-3 space-x-3">
                        <MdLowPriority size={20} />
                        <span className="font-semibold text-sm">PRIORITY</span>
                      </div>
                      <select
                        defaultValue={card.priority}
                        {...register("priority")}
                        className="rounded-md w-full"
                      >
                        <option value="NO">No</option>
                        <option value="LOW">Low</option>
                        <option value="MEDIUM">Medium</option>
                        <option value="HIGH">High</option>
                      </select>
                    </div>
                    <button
                      type="submit"
                      className="bg-indigo-600 self-end p-3 rounded text-gray-200 hover:bg-indigo-400"
                    >
                      SAVE
                    </button>
                  </form>

                  <div>
                    <div>
                      <div className="font-semibold text-sm py-3">MEMBERS</div>
                      <div className="flex flex-col space-y-5"></div>
                    </div>
                    <div>
                      <div className="font-semibold text-sm py-3">
                        MORE ACTION
                      </div>
                      <div className="flex flex-col space-y-5">
                        {moreActionBtn(
                          "DUE DATE",
                          () => {},
                          BsFillCalendarDateFill
                        )}
                        {moreActionBtn("DELETE CARD", () => {}, AiFillDelete)}
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
    </>
  );
};

export default CardDialog;
