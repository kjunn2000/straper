import { Menu, Transition } from "@headlessui/react";
import { Fragment, useEffect, useState } from "react";
import { AiFillCaretDown, AiFillDelete } from "react-icons/ai";
import { FiSettings } from "react-icons/fi";
import useIdentifyStore from "../store/identityStore";
import useWorkspaceStore from "../store/workspaceStore";
import api from "../axios/api";
import MenuItem from "./Menu/MenuItem";

export default function WorkspaceMenu() {
  const [isCreator, setIsCreator] = useState(false);
  const identity = useIdentifyStore((state) => state.identity);
  const workspace = useWorkspaceStore((state) => state.currWorkspace);
  const deleteWorkspaceAtStore = useWorkspaceStore(
    (state) => state.deleteWorkspace
  );
  const resetCurrWorkspace = useWorkspaceStore(
    (state) => state.resetCurrWorkspace
  );

  useEffect(() => {
    if (identity.user_id === workspace?.creator_id) {
      setIsCreator(true);
    } else {
      setIsCreator(false);
    }
  }, [workspace]);

  const deleteWorkspace = () => {
    api
      .post(`/protected/workspace/delete/${workspace.workspace_id}`)
      .then((res) => {
        if (res.data.Success) {
          deleteWorkspaceAtStore(workspace.workspace_id);
          resetCurrWorkspace();
        }
      });
  };

  const leaveWorkspace = () => {
    api
      .post(`/protected/workspace/leave/${workspace.workspace_id}`)
      .then((res) => {
        if (res.data.Success) {
          deleteWorkspaceAtStore(workspace.workspace_id);
          resetCurrWorkspace();
        }
      });
  };

  return (
    <div>
      <Menu as="div" className="relative w-full inline-block text-left">
        <div className="w-full" style={{ background: "rgb(47,49,54)" }}>
          <Menu.Button
            className="w-full p-3 text-sm text-white font-semibold 
          hover:color-gray-60 hover:bg-gray-600"
          >
            <div className="flex justify-between items-center">
              <span>{workspace?.workspace_name}</span>
              <AiFillCaretDown />
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
          <Menu.Items className="absolute left-0 w-56 m-5 p-2 origin-top-right bg-black divide-y divide-gray-100 rounded-md shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
            <div className="px-1 py-1">
              <MenuItem content="Workspace settings" icon={FiSettings} />
              <MenuItem
                content={isCreator ? "Delete workspace" : "Leave workspace"}
                click={isCreator ? deleteWorkspace : leaveWorkspace}
                icon={AiFillDelete}
              />
            </div>
          </Menu.Items>
        </Transition>
      </Menu>
    </div>
  );
}
