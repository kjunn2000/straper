import React from "react";
import { BsFillChatDotsFill } from "react-icons/bs";

const CardComment = () => {
  return (
    <div>
      <div className="flex self-center py-3 space-x-3">
        <BsFillChatDotsFill size={30} />
        <span className="font-semibold text-lg">ADD COMMENTS</span>
      </div>
    </div>
  );
};

export default CardComment;
