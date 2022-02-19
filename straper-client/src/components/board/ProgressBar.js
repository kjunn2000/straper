import React from "react";

const ProgressBar = ({ progressPercentage }) => {
  return (
    <div className="flex items-center space-x-3">
      <span>{progressPercentage}%</span>
      <div className="h-1 w-full bg-gray-300">
        <div
          style={{ width: `${progressPercentage}%` }}
          className={`h-full ${
            progressPercentage < 70 ? "bg-red-600" : "bg-green-600"
          }`}
        ></div>
      </div>
    </div>
  );
};

export default ProgressBar;
