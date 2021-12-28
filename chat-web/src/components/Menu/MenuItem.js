import { Menu } from "@headlessui/react";
import React from "react";
import { iconStyle } from "../../utils/style/icon";

const MenuItem = ({ content, click, icon: Icon }) => {
  return (
    <Menu.Item>
      {({ active }) => (
        <button
          className={`${
            active ? "bg-indigo-600" : ""
          } text-gray-300 font-medium group rounded-sm w-full text-sm p-3`}
          onClick={() => click()}
        >
          <div className="flex justify-between items-center">
            <Icon style={iconStyle}/>
            <span>{content}</span>
          </div>
        </button>
      )}
    </Menu.Item>
  );
};

export default MenuItem;
