import React, { useState, useEffect } from "react";
import { logOut } from "../service/logout";

const LogOutNotice = () => {
  const [timeLeft, setTimeLeft] = useState(5);

  useEffect(() => {
    const timer = setTimeout(() => {
      if (timeLeft === 0) {
        console.log("Loggin out...");
        // logOut();
        return;
      }
      setTimeLeft((count) => count - 1);
    }, 1000);
    return () => clearTimeout(timer);
  });

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
      <form className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center p-5">
        <div className="self-center">
          <div className="text-xl font-medium text-center">
            LOGOUT COUNTDOWN
          </div>
          <div className="self-center text-gray-400 text-center p-5">
            You will be logged out in
          </div>
        </div>
        <button
          type="button"
          className="bg-indigo-400 self-center w-48 p-1"
          disabled={true}
        >
          {timeLeft} second(s)
        </button>
      </form>
    </div>
  );
};

export default LogOutNotice;
