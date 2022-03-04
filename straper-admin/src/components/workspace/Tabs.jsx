import { useState } from "react";
import { Tab } from "@headlessui/react";
import classNames from "classnames";
import ChannelTable from "./ChannelTable";
import UserTable from "./UserTable";

export default function Tabs({
  userData,
  channelData,
  creatorId,
  handleRemoveUser,
  handleUpdateChannel,
  handleDeleteChannel,
}) {
  const [tables] = useState(["Users", "Channels"]);

  return (
    <div className="w-full px-2 py-16 sm:px-0">
      <Tab.Group>
        <Tab.List className="flex p-1 space-x-1 bg-blue-900/20 rounded-xl">
          {tables.map((key) => (
            <Tab
              key={key}
              className={({ selected }) =>
                classNames(
                  "w-full py-2.5 text-sm leading-5 font-medium text-blue-700 rounded-lg",
                  "focus:outline-none focus:ring-2 ring-offset-2 ring-offset-blue-400 ring-white ring-opacity-60",
                  selected
                    ? "bg-white shadow"
                    : "text-blue-100 hover:bg-white/[0.12] hover:text-white"
                )
              }
            >
              {key}
            </Tab>
          ))}
        </Tab.List>
        <Tab.Panels className="mt-2">
          {tables.map((key, idx) => (
            <Tab.Panel
              key={idx}
              className={classNames(
                "bg-white rounded-xl p-3",
                "focus:outline-none focus:ring-2 ring-offset-2 ring-offset-blue-400 ring-white ring-opacity-60"
              )}
            >
              {key === "Users" ? (
                <UserTable
                  userData={userData}
                  creatorId={creatorId}
                  handleRemoveUser={handleRemoveUser}
                />
              ) : (
                <ChannelTable
                  channelData={channelData}
                  handleUpdateChannel={handleUpdateChannel}
                  handleDeleteChannel={handleDeleteChannel}
                />
              )}
            </Tab.Panel>
          ))}
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
