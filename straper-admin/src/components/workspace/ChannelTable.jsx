import React, { useMemo, useState } from "react";
import {
  ActionCell,
  DateCell,
  IsDefaultCell,
} from "../../shared/table/TableCell";
import PaginationTable from "../../shared/table/PaginationTable";
import ActionDialog from "../../shared/dialog/ActionDialog";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import SimpleDialog from "../../shared/dialog/SimpleDialog";
import EditChannelDialog from "./EditChannelDialog";

const ChannelTable = ({
  channelData,
  handleUpdateChannel,
  handleDeleteChannel,
}) => {
  const [deleteWarningDialogOpen, setDeleteWarningDialogOpen] = useState(false);
  const [toDeleteChannelId, setToDeleteChannelId] = useState();
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");
  const [editChannel, setEditChannel] = useState();
  const history = useHistory();

  const getChannel = (channelId) => {
    return channelData.find((channel) => channel.channel_id === channelId);
  };

  const isDefaultChannel = (channelId) => {
    return getChannel(channelId)?.is_default;
  };

  const columns = useMemo(
    () => [
      {
        Header: "ID",
        accessor: "channel_id",
      },
      {
        Header: "Name",
        accessor: "channel_name",
      },
      {
        Header: "Creator ID",
        accessor: "creator_id",
      },
      {
        Header: "Created",
        accessor: "created_date",
        Cell: DateCell,
      },
      {
        Header: "Default",
        accessor: "is_default",
        Cell: IsDefaultCell,
      },
      {
        Header: "Action",
        idAccessor: "channel_id",
        Cell: ActionCell,
        editAction: (channelId) => {
          setEditChannel(getChannel(channelId));
          setShowEditDialog(true);
        },
        deleteAction: (channelId) => {
          if (isDefaultChannel(channelId)) {
            setDialogErrMsg("Cannot delete default channel.");
            setShowFailDialog(true);
            return;
          }
          setToDeleteChannelId(channelId);
          setDeleteWarningDialogOpen(true);
        },
      },
    ],
    []
  );

  const data = useMemo(() => channelData);

  return (
    <div>
      <PaginationTable columns={columns} data={data} />

      <ActionDialog
        isOpen={deleteWarningDialogOpen}
        setIsOpen={setDeleteWarningDialogOpen}
        title="Delete Channel Confirmation"
        content="Please confirm that the deleted channel will not able to be recovered."
        buttonText="Delete Anyway"
        buttonStatus="fail"
        buttonAction={() => handleDeleteChannel(toDeleteChannelId)}
        closeButtonText="Close"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Update Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />

      <EditChannelDialog
        isOpen={showEditDialog}
        close={() => setShowEditDialog(false)}
        handleUpdateChannel={handleUpdateChannel}
        channel={editChannel}
      />
    </div>
  );
};

export default ChannelTable;
