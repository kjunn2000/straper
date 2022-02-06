import { Menu, Transition } from "@headlessui/react";
import { Fragment, useEffect, useState } from "react";
import { AiFillEdit, AiFillDelete, AiOutlineLink } from "react-icons/ai";
import { FiSettings, FiMoreHorizontal } from "react-icons/fi";
import useIdentityStore from "../../store/identityStore";
import useWorkspaceStore from "../../store/workspaceStore";
import api from "../../axios/api";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import ActionDialog from "../dialog/ActionDialog";
import SimpleDialog from "../dialog/SimpleDialog";
import { copyTextToClipboard } from "../../service/navigator";
import MenuItem from "../menu/MenuItem";

export default function MessageMenu({ type, editMsg, deleteMsg }) {
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);

  return (
    <div>
      <Menu as="div" className="relative w-full inline-block text-left">
        <div className="w-full">
          <Menu.Button className="w-full p-3 text-sm text-white">
            <div className="flex justify-between items-center">
              <FiMoreHorizontal
                size={18}
                className="opacity-0 group-hover:opacity-100 trasition duration-300 text-gray-400"
              />
            </div>
          </Menu.Button>
        </div>
        <Transition
          as={Fragment}
          enter="transition ease-out duration-100"
          enterFrom="transform opacity-0 scale-95"
          enterTo="transform opacity-100 scale-100"
          leave="transition ease-in duration-75"
          leaveFrom="transform opacity-100 scale-100"
          leaveTo="transform opacity-0 scale-95"
        >
          <Menu.Items className="absolute right-0 w-56 m-5 p-2 origin-top-right bg-black divide-y divide-gray-100 rounded-md shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
            <div className="px-1 py-1">
              {type === "MESSAGE" && (
                <MenuItem
                  content="Edit"
                  icon={AiFillEdit}
                  click={() => {
                    editMsg();
                  }}
                />
              )}
              <MenuItem
                content="Delete"
                click={() => setDeleteWarningDialogOpen(true)}
                icon={AiFillDelete}
              />
            </div>
          </Menu.Items>
        </Transition>
      </Menu>
      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Delete Message Confirmation"
        content="Please confirm that the deleted message will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={() => deleteMsg()}
        closeButtonText="Close"
      />
    </div>
  );
}
