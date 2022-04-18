import { Popover } from "@headlessui/react";
import { AiOutlineCheckCircle } from "react-icons/ai";
import { FaDotCircle } from "react-icons/fa";
import { iconStyle } from "../../utils/style/icon";

export default function AccountPopOver({ children, user, currStatus }) {
  return (
    <Popover>
      <Popover.Button className="w-full">{children}</Popover.Button>

      <Popover.Panel className="absolute z-10 m-3 -translate-x-1/2">
        <div className="bg-black text-white rounded-lg flex flex-col">
          <div className="bg-indigo-500 p-5 rounded-lg">
            <span className="text-2xl font-bold">{user.username}</span>
            <span className="text-gray-200 font-semibold p-1">
              #{user.user_id}
            </span>
          </div>
          <div className="flex flex-col space-y-4 space-x-1 p-5 rounded-lg">
            <div className="flex justify-end items-center text-right text-gray-500">
              {currStatus}
              {currStatus === "Active" ? (
                <AiOutlineCheckCircle
                  style={iconStyle}
                  className="text-green-400"
                />
              ) : (
                <FaDotCircle style={iconStyle} className="text-gray-600" />
              )}
            </div>
            <div className="flex flex-col">
              <span className="font-semibold">EMAIL</span>
              <span className="text-gray-500 font-semibold">{user.email}</span>
            </div>
            <div className="flex flex-col">
              <span className="font-semibold">PHONE NO</span>
              <span className="text-gray-500 font-semibold">
                {user.phone_no}
              </span>
            </div>
          </div>
        </div>
      </Popover.Panel>
    </Popover>
  );
}
