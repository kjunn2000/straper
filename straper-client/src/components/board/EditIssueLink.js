import React, { useEffect, useState } from "react";
import { Popover } from "@headlessui/react";
import { AiFillEdit } from "react-icons/ai";
import Select from "react-select";
import { sendBoardMsg } from "../../service/websocket";
import useIssueStore from "../../store/issueStore";
import { getIssueData } from "../../service/bug";
import useBoardStore from "../../store/boardStore";
import SimpleDialog from "../../shared/dialog/SimpleDialog";
import { isEmpty, isObjectEmpty } from "../../service/object";

const EditIssueLink = ({ card }) => {
  const isSet = useIssueStore((state) => state.isSet);
  const issues = useIssueStore((state) => state.issues);
  const board = useBoardStore((state) => state.board);
  const [options, setOptions] = useState([]);
  const [issueLink, setIssueLink] = useState({});
  const [successEditOpen, setSuccessEditOpen] = useState(false);

  useEffect(() => {
    if (!isSet) {
      getIssueData();
    }
    setOptions(convertToOptions(issues));
  }, [issues]);

  const convertToOptions = (issues) => {
    const options = issues.map((issue) => ({
      value: issue.issue_id,
      label: issue.summary,
    }));
    options.push({ value: undefined, label: "No" });
    return options;
  };

  const onEdit = (e, close) => {
    e.preventDefault();
    if (!issueLink || isObjectEmpty(issueLink)) {
      return;
    }
    const payload = {
      list_id: card.list_id,
      card_id: card.card_id,
      issue_link: issueLink.value,
    };
    sendBoardMsg("BOARD_UPDATE_CARD_ISSUE_LINK", board.workspace_id, payload);
    setIssueLink({});
    close();
    setSuccessEditOpen(true);
  };

  return (
    <Popover className="relative z-30">
      <Popover.Button>
        <AiFillEdit />
      </Popover.Button>

      <Popover.Panel className="absolute w-screen max-w-xs px-4 mt-3 transform -translate-x-1/2 -translate-y-1/2 left-1/2 sm:px-0 md:max-w-xl lg:max-w-3xl">
        {({ close }) => (
          <form
            className="bg-white rounded-xl flex flex-col space-y-1 w-72 md:w-96 h-auto justify-center self-center py-5 px-3"
            onSubmit={(e) => onEdit(e, close)}
          >
            <div className="self-center w-full">
              <div className="text-xl font-medium text-center ">ISSUE LINK</div>
            </div>
            {options && options.length > 0 ? (
              <div>
                <div className="self-center w-full p-3">
                  <Select
                    options={options}
                    autoFocus={true}
                    onChange={(e) => setIssueLink(e)}
                  />
                </div>
                <button
                  type="submit"
                  className="self-end bg-indigo-400 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded-full"
                >
                  EDIT
                </button>
              </div>
            ) : (
              <span className="text-center font-semibold text-sm">
                NO ISSUE AVAILABLE
              </span>
            )}
          </form>
        )}
      </Popover.Panel>
      <SimpleDialog
        isOpen={successEditOpen}
        setIsOpen={setSuccessEditOpen}
        title="Success Edit"
        content="Successfully edited the issue link of this card."
        buttonText="Close"
        buttonStatus="success"
      />
    </Popover>
  );
};

export default EditIssueLink;
