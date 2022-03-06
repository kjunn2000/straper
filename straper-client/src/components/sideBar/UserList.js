import React, { useEffect, useState } from "react";
import useIdentityStore from "../../store/identityStore";
import useWorkspaceStore from "../../store/workspaceStore";
import { iconStyle } from "../../utils/style/icon";
import { AiOutlineCheckCircle } from "react-icons/ai";
import AccountPopOver from "./AccountPopOver";

const UserList = () => {
  const currAccountList = useWorkspaceStore((state) => state.currAccountList);
  const [accountList, setAccountList] = useState({
    activeList: [],
    offlineList: [],
  });
  const identity = useIdentityStore((state) => state.identity);

  useEffect(() => {
    const currTime = new Date();
    const parsedAccountList = Object.values(currAccountList).map((val) => {
      const lastSeen = new Date(val.last_seen);
      val.last_seen_minute =
        identity.user_id === val.user_id ? 0 : calTimeDiff(currTime, lastSeen);
      return val;
    });
    const activeList = parsedAccountList.filter(
      (user) => user.last_seen_minute < 5
    );
    const offlineList = parsedAccountList.filter(
      (user) => user.last_seen_minute >= 5
    );
    setAccountList({ activeList, offlineList });
  }, [currAccountList]);

  const calTimeDiff = (dt1, dt2) => {
    var diff = (dt2.getTime() - dt1.getTime()) / 1000;
    diff /= 60;
    return Math.abs(Math.round(diff));
  };

  const parseMinuteDiff = (minute) => {
    if (minute < 5) {
      return "Active";
    } else if (minute < 60) {
      return minute + " minutes ago";
    } else if (minute / 60 < 24) {
      return Math.round(minute / 60) + " hours ago";
    } else if (minute / 60 < 48) {
      return "One day ago";
    } else {
      return "Few days ago";
    }
  };

  const userRow = (user) => (
    <AccountPopOver
      user={user}
      currStatus={parseMinuteDiff(user.last_seen_minute)}
    >
      <div
        className="group flex text-gray-400 text-lg items-center 
     hover:bg-gray-700 rounded-xl hover:cursor-pointer p-1"
      >
        <div className="w-1/5">
          <img
            src={generateDiceBearBottts(Math.random())}
            width="40"
            alt="bottts_avatar"
          />
        </div>
        <span className="w-2/5">{user.username}</span>
        <span className="w-2/5 flex items-center opacity-0 group-hover:opacity-100 text-gray-500 text-xs">
          {parseMinuteDiff(user.last_seen_minute)}
          {user.last_seen_minute < 5 && (
            <AiOutlineCheckCircle
              style={iconStyle}
              className="text-green-400"
            />
          )}
        </span>
      </div>
    </AccountPopOver>
  );
  const generateDiceBearBottts = (seed) =>
    `https://avatars.dicebear.com/api/bottts/${seed}.svg`;

  return (
    <div className="flex flex-col space-y-3 p-3 overflow-y-auto h-full">
      <div>
        <div className="text-gray-400 font-semibold text-sm">
          ONLINE - {accountList.activeList.length}
        </div>
        <div className="flex flex-col space-y-2 py-3">
          {accountList.activeList.map((user) => userRow(user))}
        </div>
      </div>
      <div>
        <div className="text-gray-400 font-semibold text-sm">
          OFFLINE - {accountList.offlineList.length}
        </div>
        <div className="flex flex-col space-y-2 py-3">
          {accountList.offlineList.map((user) => userRow(user))}
        </div>
      </div>
    </div>
  );
};

export default UserList;
