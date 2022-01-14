import React from "react";

const SidebarIcon = ({ content, click, bgColor, hoverBgColor }) => {
  return (
    <div className="w-auto text-center py-3">
      <button
        className={`relative flex items-center justify-center 
               h-12 w-12 mt-2 mb-2 mx-auto shadow-lg
               bg-gray-800 text-white text-white rounded-3xl hover:rounded-xl
               transition-all duration-300 ease-linear
               cursor-pointer
	 ${bgColor} ${hoverBgColor}`}
        onClick={() => click()}
      >
        {content.toUpperCase().substring(0, 2)}
      </button>
    </div>
  );
};

export default SidebarIcon;
