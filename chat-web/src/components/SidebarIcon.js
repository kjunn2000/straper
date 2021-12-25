import React from "react";

const SidebarIcon = ({ content, click, bgColor, hoverBgColor }) => {
  return (
    <div className="w-auto text-center py-3">
      <button
        className={`rounded-full hover:rounded text-white text-center h-12 w-12 items-center justify-center font-semibold
	 ${bgColor} ${hoverBgColor}`}
        onClick={() => click()}
      >
        {content.toUpperCase().substring(0, 2)}
      </button>
    </div>
  );
};

export default SidebarIcon;
