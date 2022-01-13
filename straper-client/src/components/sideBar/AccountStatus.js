import React from "react";
import { FaUserCircle } from "react-icons/fa";
import { FiSettings } from "react-icons/fi";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import useIdentityStore from "../../store/identityStore";

const AccountStatus = () => {
  const identity = useIdentityStore((state) => state.identity);
  const history = useHistory();

  return (
    <div
      className="flex shadow-2xl bg-neutral-700 rounded px-5 justify-between
      items-center text-sm text-gray-400 hover:text-white transition-all duration-200"
      style={{ background: "rgb(32,34,37)" }}
    >
      <FaUserCircle size={20} />
      <div className="px-5">
        <div>{identity.username}</div>
        <div>{identity.user_id}</div>
      </div>
      <FiSettings
        className="hover:cursor-pointer"
        size={20}
        onClick={() => {
          history.push("/setting");
        }}
      />
    </div>
  );
};

export default AccountStatus;
