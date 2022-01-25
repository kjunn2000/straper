import React, { useEffect, useRef, useState } from "react";
import { Popover } from "@headlessui/react";
import { BsPlusLg } from "react-icons/bs";
import Select from "react-select";
import api from "../../axios/api";
import { isEmpty } from "../../service/object";
import useWorkspaceStore from "../../store/workspaceStore";

const AddMember = ({ card }) => {
  const [memberOptions, setMemberOptions] = useState([]);
  const [toAddMembers, setToAddMembers] = useState([]);
  const currAccountList = useWorkspaceStore((state) => state.currAccountList);

  useEffect(() => {
    getCardMember();
  }, []);

  const getCardMember = () => {
    api.get(`/protected/board/card/member/${card.card_id}`).then((res) => {
      if (!res.data.Success) {
        return;
      }
      const newAccountList = currAccountList;
      if (!isEmpty(res.data.Data)) {
        res.data.Data.foreach((userId) => {
          delete newAccountList[userId];
        });
      }
      const newMemberOptions = Object.entries(newAccountList).map((entry) => ({
        value: entry[0],
        label: entry[1].username,
      }));
      setMemberOptions(newMemberOptions);
    });
  };

  const handleAddMember = (e) => {
    e.preventDefault();
    const payload = toAddMembers.map((member) => member.value);
    console.log(payload);
  };

  return (
    <Popover className="relative">
      <Popover.Button>
        <BsPlusLg />
      </Popover.Button>

      <Popover.Panel className="absolute w-screen max-w-sm px-4 mt-3 transform -translate-x-1/2 left-1/2 sm:px-0 lg:max-w-3xl">
        <form
          className="bg-white rounded-xl flex flex-col space-y-1 w-96 h-auto justify-center self-center py-5 px-3"
          onSubmit={(e) => handleAddMember(e)}
        >
          <div className="self-center w-full">
            <div className="text-xl font-medium text-center ">MEMBER</div>
          </div>
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
        </form>
      </Popover.Panel>
    </Popover>
  );
};

export default AddMember;
