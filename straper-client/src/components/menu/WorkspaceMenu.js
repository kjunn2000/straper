import { Menu, Transition } from "@headlessui/react";
import { Fragment, useEffect, useState } from "react";
import { AiFillCaretDown, AiFillDelete, AiOutlineLink } from "react-icons/ai";
import { FiSettings } from "react-icons/fi";
import useIdentityStore from "../../store/identityStore";
import useWorkspaceStore from "../../store/workspaceStore";
import api from "../../axios/api";
import MenuItem from "./MenuItem";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import { copyTextToClipboard } from "../../service/navigator";
import ActionDialog from "../../shared/dialog/ActionDialog";
import SimpleDialog from "../../shared/dialog/SimpleDialog";

export default function WorkspaceMenu() {
  const [isCreator, setIsCreator] = useState(false);
  const identity = useIdentityStore((state) => state.identity);
  const workspace = useWorkspaceStore((state) => state.currWorkspace);
  const deleteWorkspaceAtStore = useWorkspaceStore(
    (state) => state.deleteWorkspace
  );
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const selectedChannelIds = useWorkspaceStore(
    (state) => state.selectedChannelIds
  );
  const deleteSelectedChannelIds = useWorkspaceStore(
    (state) => state.deleteSelectedChannelIds
  );
  const clearWorkspaceState = useWorkspaceStore(
    (state) => state.clearWorkspaceState
  );
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [successCopyLinkDialogOpen, setSuccessCopyLinkDialogOpen] =
    useState(false);

  const history = useHistory();

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
          updateWorkspaceState();
        }
      });
  };

  const leaveWorkspace = () => {
    api
      .post(`/protected/workspace/leave/${workspace.workspace_id}`)
      .then((res) => {
        if (res.data.Success) {
          updateWorkspaceState();
        }
      });
  };

  const updateWorkspaceState = () => {
    deleteWorkspaceAtStore(workspace.workspace_id);
    deleteSelectedChannelIds(workspace.workspace_id);
    const selectedIds = [...selectedChannelIds];
    if (selectedIds.length > 0) {
      setCurrWorkspace(selectedIds[0][0]);
      setCurrChannel(selectedIds[0][1]);
      history.push(`/channel/${selectedIds[0][0]}/${selectedIds[0][1]}`);
    } else {
      clearWorkspaceState();
      history.push("/channel");
    }
  };

  const copyLinkToClipboard = () => {
    copyTextToClipboard(workspace.workspace_id);
    setSuccessCopyLinkDialogOpen(true);
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
                content="Invitation link"
                icon={AiOutlineLink}
                click={() => {
                  copyLinkToClipboard();
                }}
              />
              <MenuItem
                content={isCreator ? "Delete workspace" : "Leave workspace"}
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
        title="Delete Workspace Confirmation"
        content="Please confirm that the deleted workspace will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={isCreator ? deleteWorkspace : leaveWorkspace}
        closeButtonText="Close"
      />
      <SimpleDialog
        isOpen={successCopyLinkDialogOpen}
        setIsOpen={setSuccessCopyLinkDialogOpen}
        title="Copied Workspace Link To Clipboard"
        content="Successfully copied workspace link to clipboard. You are able to send it to your friend for joining this workspace."
        buttonText="Close"
        buttonStatus="success"
      />
    </div>
  );
}
