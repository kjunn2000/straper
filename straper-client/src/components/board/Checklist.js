import React, { useState } from "react";
import { BsCardChecklist } from "react-icons/bs";
import ProgressBar from "./ProgressBar";

const Checklist = ({ show }) => {
  const [percentage, setPercentage] = useState(100);
  return (
    <div className={show ? "" : "hidden"}>
      <div className="flex self-center py-3 space-x-3">
        <BsCardChecklist size={30} />
        <span className="font-semibold text-lg">TO DO's</span>
      </div>
      <ProgressBar progressPercentage={percentage} />
      <div className="text-gray-500 text-sm hover:cursor-pointer">
        Add an item
      </div>
    </div>
  );
};

export default Checklist;
