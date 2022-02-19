import React, { useEffect, useState } from "react";
import { Popover } from "@headlessui/react";
import { BsPlusLg } from "react-icons/bs";
import Select from "react-select";
import useWorkspaceStore from "../../store/workspaceStore";
import { sendBoardMsg } from "../../service/websocket";
import useBoardStore from "../../store/boardStore";

const AddMember = ({ card }) => {
  const [memberOptions, setMemberOptions] = useState([]);
  const [toAddMembers, setToAddMembers] = useState([]);
  const board = useBoardStore((state) => state.board);
  const currAccountList = useWorkspaceStore((state) => state.currAccountList);

  useEffect(() => {
    if (!card.member_list) {
      setMemberOptions(convertAccountListToOptions(currAccountList));
      return;
    }
    const newMemberOptions = JSON.parse(JSON.stringify(currAccountList));
    card.member_list.forEach((userId) => {
      delete newMemberOptions[userId];
    });
    setMemberOptions(convertAccountListToOptions(newMemberOptions));
    setToAddMembers([]);
  }, [card.member_list, currAccountList]);

  const convertAccountListToOptions = (accountList) => {
    return Object.entries(accountList).map((entry) => ({
      value: entry[0],
      label: entry[1].username,
    }));
  };

  const handleAddMember = (e, close) => {
    e.preventDefault();
    if (!toAddMembers || toAddMembers.length === 0) {
      return;
    }
    const payload = {
      list_id: card.list_id,
      card_id: card.card_id,
      member_list: toAddMembers.map((member) => member.value),
    };
    sendBoardMsg("BOARD_CARD_ADD_MEMBERS", board.workspace_id, payload);
    close();
  };

  return (
    <Popover className="relative">
      <Popover.Button>
        <BsPlusLg />
      </Popover.Button>

      <Popover.Panel className="absolute w-screen max-w-sm px-4 mt-3 transform -translate-x-1/2 left-1/2 sm:px-0 lg:max-w-3xl">
        {({ close }) => (
          <form
            className="bg-white rounded-xl flex flex-col space-y-1 w-96 h-auto justify-center self-center py-5 px-3"
            onSubmit={(e) => handleAddMember(e, close)}
          >
            <div className="self-center w-full">
              <div className="text-xl font-medium text-center ">MEMBER</div>
            </div>
            {memberOptions && memberOptions.length > 0 ? (
              <div>
                <div className="self-center w-full p-3">
                  <Select
                    isMulti={true}
                    options={memberOptions}
                    autoFocus={true}
                    onChange={(e) => setToAddMembers(e)}
                  />
                </div>
                <button
                  type="submit"
                  className="self-end bg-indigo-400 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded-full"
                >
                  ADD
                </button>
              </div>
            ) : (
              <span className="text-center font-semibold text-sm">
                NO MEMBER AVAILABLE
              </span>
            )}
          </form>
        )}
      </Popover.Panel>
    </Popover>
  );
};

export default AddMember;
