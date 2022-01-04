import axios from "axios";
import React, { useEffect, useState } from "react";
import { useHistory } from "react-router";

const EmailVerify = () => {
  const history = useHistory();
  const [verifySuccess, setVerifySuccess] = useState(false);

  useEffect(() => {
    const tokenId = history.location.pathname.split("/").pop();
    axios
      .post(
        `http://localhost:8080/api/v1/account/email/verify/${tokenId}`,
        tokenId
      )
      .then((res) => {
        if (res.data.Success) {
          setVerifySuccess(true);
        } else {
          setVerifySuccess(false);
        }
      })
      .catch(() => {
        setVerifySuccess(false);
      });
  }, []);

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
      <form className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center p-5">
        <div className="self-center">
          <div className="text-xl font-medium text-center">
            {verifySuccess ? "EMAIL VERIFIED" : "EMAIL VERIFICATION FAILED"}
          </div>
          <div className="self-center text-gray-400 text-center p-5">
            {verifySuccess
              ? "Thank you for your verifying, let enjoy with the power of Straper."
              : "Something went wrong, email verification record is not exist."}
          </div>
        </div>
        <button
          type="button"
          className="bg-indigo-400 self-center w-48 p-1"
          onClick={() => history.push("/login")}
        >
          Login Straper
        </button>
      </form>
    </div>
  );
};

export default EmailVerify;
