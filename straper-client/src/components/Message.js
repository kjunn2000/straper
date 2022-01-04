import React from "react";
import useIdentifyStore from "../store/identityStore";

const Message = ({ msg }) => {
  const identity = useIdentifyStore((state) => state.identity);

  const isCreator = () => {
    return identity.username == msg.creator_name;
  };

  const convertToDateString = (timestamp) => {
    var date = new Date(timestamp);
    var dd = date.getDate();
    var mm = date.getMonth() + 1;
    var yy = date.getFullYear();
    var hour = date.getHours();
    var min = date.getMinutes();
    return dd + "/" + mm + "/" + yy + " " + hour + ":" + min;
  };

  return (
    <div className="chat-message">
      <div
        className={`flex ${
          isCreator() ? "items-end justify-end" : "items-start justify-start"
        }`}
      >
        {isCreator() ? (
          <div className="flex flex-col max-w-xs mx-2 items-end">
            <span className="inline-block text-gray-300">
              {msg.creator_name}
            </span>
            <span className="rounded-lg inline-block rounded-br-none bg-indigo-500 text-white ">
              {msg.content}
            </span>
            <span>{convertToDateString(msg.created_date)}</span>
          </div>
        ) : (
          <div className="flex flex-col max-w-xs mx-2 items-start">
            <span className="inline-block text-gray-300">
              {msg.creator_name}
            </span>
            <span className="rounded-lg inline-block rounded-br-none bg-sky-400 text-white ">
              {msg.content}
            </span>
            <span>{convertToDateString(msg.created_date)}</span>
          </div>
        )}
      </div>
    </div>
  );
};

export default Message;
