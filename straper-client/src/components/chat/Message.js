import React from "react";
import useIdentifyStore from "../../store/identityStore";

const Message = ({ msg }) => {
  const identity = useIdentifyStore((state) => state.identity);

  const isCreator = () => {
    return identity.username === msg.creator_name;
  };

  const zeroPad = (num, places) => String(num).padStart(places, "0");

  const convertToDateString = (timestamp) => {
    var date = new Date(timestamp);
    var dd = date.getDate();
    var mm = date.getMonth() + 1;
    var yy = date.getFullYear();
    var hour = date.getHours();
    var min = date.getMinutes();
    return (
      dd + "/" + mm + "/" + yy + " " + zeroPad(hour, 2) + ":" + zeroPad(min, 2)
    );
  };

  return (
    <div className="chat-message">
      <div
        className={`flex ${
          isCreator() ? "items-end justify-end" : "items-start justify-start"
        }`}
      >
        <div
          className={`flex flex-col max-w-xs mx-2 group ${
            isCreator() ? "items-end" : "items-start"
          }`}
        >
          <span className="inline-block text-gray-300">{msg.creator_name}</span>
          <p
            className={`px-3 py-2 rounded-lg inline-block text-white max-w-sm break-words ${
              isCreator()
                ? "rounded-br-none bg-indigo-500"
                : "rounded-bl-none bg-gray-500"
            }`}
          >
            {msg.content}
          </p>
          <span className="invisible text-gray-400 group-hover:visible transition duration-150">
            {convertToDateString(msg.created_date)}
          </span>
        </div>
      </div>
    </div>
  );
};

export default Message;
