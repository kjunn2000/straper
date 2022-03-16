import React, { useEffect, useState } from "react";
import { AiFillDelete, AiOutlineLink, AiOutlinePlus } from "react-icons/ai";
import { BsDoorOpen } from "react-icons/bs";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import api from "../../axios/api";
import { copyTextToClipboard } from "../../service/navigator";
import useIdentityStore from "../../store/identityStore";
import useWorkspaceStore from "../../store/workspaceStore";
import { darkGrayBg } from "../../utils/style/color";
import { iconStyle } from "../../utils/style/icon";
import ActionDialog from "../../shared/dialog/ActionDialog";
import AddDialog from "./AddDialog";
import JoinDialog from "./JoinDialog";
import SimpleDialog from "../../shared/dialog/SimpleDialog";
import WorkspaceMenu from "../menu/WorkspaceMenu";
import AccountStatus from "./AccountStatus";

function WorkspaceSidebar() {
  const history = useHistory();

  const identity = useIdentityStore((state) => state.identity);
  const currWorkspace = useWorkspaceStore((state) => state.currWorkspace);
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const setCurrWorkspace = useWorkspaceStore((state) => state.setCurrWorkspace);
  const setCurrChannel = useWorkspaceStore((state) => state.setCurrChannel);
  const setSelectedChannelIds = useWorkspaceStore(
    (state) => state.setSelectedChannelIds
  );
  const deleteChannelFromWorkspace = useWorkspaceStore(
    (state) => state.deleteChannelFromWorkspace
  );
  const selectedChannelIds = useWorkspaceStore(
    (state) => state.selectedChannelIds
  );
  const addChannelToWorkspace = useWorkspaceStore(
    (state) => state.addChannelToWorkspace
  );

  const [failDeleteDialogOpen, setFailDeleteDialogOpen] = useState(false);
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [targetDeleteChannelId, setTargetDeleteChannelId] = useState("");
  const [deleteType, setDeleteType] = useState("");
  const [isAddChannelDialogOpen, setAddChannelDialogOpen] = useState(false);
  const [isJoinChannelDialogOpen, setJoinChannelDialogOpen] = useState(false);
  const [successCopyLinkDialogOpen, setSuccessCopyLinkDialogOpen] =
    useState(false);
  const [invalidChannelDialogOpen, setInvalidChannelDialogOpen] =
    useState(false);

  const changeChannel = (channelId) => {
    if (channelId === currChannel.channel_id) {
      return;
    }
    setCurrChannel(channelId);
    setSelectedChannelIds(currWorkspace.workspace_id, channelId);
    history.push(`/channel/${currWorkspace.workspace_id}/${channelId}`);
  };

  const onDeleteChannel = (e, channelId, type) => {
    e.stopPropagation();
    if (currWorkspace.channel_list.length === 1) {
      setFailDeleteDialogOpen(true);
      return;
    }
    setTargetDeleteChannelId(channelId);
    setDeleteType(type);
    setDeleteWarningDialogOpen(true);
  };

  const deleteOrLeaveChannel = async (channelId) => {
    const res = await api.post(`/protected/channel/${deleteType}/${channelId}`);
    if (res.data.Success) {
      deleteChannelFromWorkspace(channelId);
      setCurrWorkspace(currWorkspace.workspace_id);
      if (selectedChannelIds.get(currWorkspace.workspace_id) === channelId) {
        const nextChannelId = currWorkspace.channel_list[0].channel_id;
        setCurrChannel(nextChannelId);
        setSelectedChannelIds(currWorkspace.workspace_id, nextChannelId);
        history.push(`/channel/${currWorkspace.workspace_id}/${nextChannelId}`);
      }
    } else {
      setInvalidChannelDialogOpen(true);
    }
  };

  const toggleDialog = () => {
    if (isAddChannelDialogOpen) {
      setAddChannelDialogOpen(false);
      setJoinChannelDialogOpen(true);
    } else {
      setJoinChannelDialogOpen(false);
      setAddChannelDialogOpen(true);
    }
  };

  const addNewChannel = async (data) => {
    const dto = {
      workspace_id: currWorkspace.workspace_id,
      channel_name: data?.channel_name,
    };

    const res = await api.post("/protected/channel/create", dto);
    if (res.data.Success) {
      updateNewChannel(res.data.Data);
      return;
    } else {
      switch (res.data.ErrorMessage) {
        case "workspace.id.not.found": {
          return "Workspace may be deleted, please refresh the page.";
        }
      }
    }
  };

  const joinNewChannel = async (data) => {
    const dto = {
      workspace_id: currWorkspace.workspace_id,
      channel_id: data?.channel_id,
    };

    const res = await api.post("/protected/channel/join", dto);
    if (res.data.Success) {
      updateNewChannel(res.data.Data);
      return;
    } else {
      switch (res.data.ErrorMessage) {
        case "channel.user.record.exist": {
          return "You has been joined to this channel.";
        }
        case "workspace.id.not.found": {
          return "Workspace may be deleted, please refresh the page.";
        }
        case "channel.id.not.found": {
          return "Invalid channel id.";
        }
      }
    }
  };

  const updateNewChannel = (channel) => {
    addChannelToWorkspace(currWorkspace.workspace_id, channel);
    setCurrWorkspace(currWorkspace.workspace_id);
    setCurrChannel(channel.channel_id);
    setSelectedChannelIds(currWorkspace.workspace_id, channel.channel_id);
    history.push(
      `/channel/${currWorkspace.workspace_id}/${channel.channel_id}`
    );
  };

  const copyLinkToClipboard = (e, channelId) => {
    e.stopPropagation();
    copyTextToClipboard(channelId);
    setSuccessCopyLinkDialogOpen(true);
  };

  const channelList =
    currWorkspace?.channel_list &&
    currWorkspace.channel_list.map((channel) => (
      <div
        className="group flex hover:cursor-pointer justify-between text-white text-sm font-medium p-3 text-gray-400 hover:bg-gray-700 rounded hover:text-white"
        key={channel?.channel_id}
        onClick={() => changeChannel(channel.channel_id)}
      >
        <span> # {channel?.channel_name} </span>
        <div className="flex">
          <span
            className="opacity-0 group-hover:opacity-100 cursor-pointer"
            onClick={(e) => copyLinkToClipboard(e, channel.channel_id)}
          >
            <AiOutlineLink style={iconStyle} />
          </span>
          {!channel.is_default &&
            (identity.user_id === channel.creator_id ? (
              <span
                className="opacity-0 group-hover:opacity-100 cursor-pointer pl-3"
                onClick={(e) =>
                  onDeleteChannel(e, channel.channel_id, "delete")
                }
              >
                <AiFillDelete style={iconStyle} />
              </span>
            ) : (
              <span
                className="opacity-0 group-hover:opacity-100 cursor-pointer pl-3"
                onClick={(e) => onDeleteChannel(e, channel.channel_id, "leave")}
              >
                <BsDoorOpen style={iconStyle} />
              </span>
            ))}
        </div>
      </div>
    ));

  return (
    <div className="flex flex-col justify-between w-full" style={darkGrayBg}>
      {currWorkspace.workspace_id ? (
        <div className="flex flex-col">
          <WorkspaceMenu />
          <div
            className="p-3 text-sm text-gray-400 cursor-pointer hover:text-white"
            onClick={() => history.push("/board")}
          >
            <span>TASK BOARD</span>
          </div>
          <div
            className="p-3 text-sm text-gray-400 cursor-pointer hover:text-white"
            onClick={() => history.push("/bug")}
          >
            <span>BUG DASHBOARD</span>
          </div>
          <div>
            <div className="p-3 text-sm text-gray-400 flex justify-between hover:text-white">
              <span>CHANNELS</span>
              <AiOutlinePlus onClick={() => setAddChannelDialogOpen(true)} />
            </div>
            <div className="px-3">{channelList}</div>
          </div>
          <AddDialog
            isOpen={isAddChannelDialogOpen}
            close={() => setAddChannelDialogOpen(false)}
            toggleDialog={toggleDialog}
            addAction={addNewChannel}
            type="channel"
          />
          <JoinDialog
            isOpen={isJoinChannelDialogOpen}
            close={() => setJoinChannelDialogOpen(false)}
            toggleDialog={toggleDialog}
            joinAction={joinNewChannel}
            type="channel"
          />
          <SimpleDialog
            isOpen={failDeleteDialogOpen}
            setIsOpen={setFailDeleteDialogOpen}
            title="Fail To Delete Channel"
            content="Please note that you are not allowed to delete/leave default channel."
            buttonText="Close"
            buttonStatus="fail"
          />
          <SimpleDialog
            isOpen={invalidChannelDialogOpen}
            setIsOpen={setInvalidChannelDialogOpen}
            title="Channel Not Found"
            content="The workspace may be deleted by the creator, please refresh your page."
            buttonText="Close"
            buttonStatus="fail"
          />
          <ActionDialog
            isOpen={deleteWarningDialogOpen}
            setIsOpen={setDeleteWarningDialogOpen}
            title={
              deleteType === "delete"
                ? "Delete Channel Confirmation"
                : "Leave Channel Confirmation"
            }
            content="Please confirm that the removed channel will not able to be recovered."
            buttonText={
              deleteType === "delete" ? "Delete Anyway" : "Leave Anyway"
            }
            buttonStatus="fail"
            buttonAction={() => deleteOrLeaveChannel(targetDeleteChannelId)}
            closeButtonText="close"
          />
          <SimpleDialog
            isOpen={successCopyLinkDialogOpen}
            setIsOpen={setSuccessCopyLinkDialogOpen}
            title="Copied Channel Link To Clipboard"
            content="Successfully copied channel link to clipboard. You are able to send it to your friend for joining this channel."
            buttonText="Close"
            buttonStatus="success"
          />
        </div>
      ) : (
        <div className="flex flex-col"></div>
      )}
      <div className="text-white w-full">
        <AccountStatus />
      </div>
    </div>
  );
}

export default WorkspaceSidebar;
