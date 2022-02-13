import React, { useState } from "react";
import { useForm } from "react-hook-form";
import api from "../../axios/api";
import useIdentityStore from "../../store/identityStore";
import SimpleDialog from "../dialog/SimpleDialog";

const UserPassword = () => {
  const {
    handleSubmit,
  } = useForm();
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");
  const identity = useIdentityStore((state) => state.identity);

  const onReset = () => {
    const payload = {
      email: identity.email,
    };
    api
      .post(
        "/account/reset-password/create",
        payload
      )
      .then((res) => {
        if (res.data.Success) {
          setShowSuccessDialog(true);
        } else {
          switch (res.data.ErrorMessage) {
            case "email.not.found": {
              setDialogErrMsg("Email not found. Please try again.");
              break;
            }
            case "invalid.email.format": {
              setDialogErrMsg("Email format incorrect. Please try again.");
              break;
            }
            case "password_reset_attempt_in_past_15_min": {
              setDialogErrMsg(
                "Password reset request has been sent to your email inbox in the past 15 minutes. Please check it out."
              );
              break;
            }
            default: {
              setDialogErrMsg("Something went wrong. Please try again.");
            }
          }
          setShowFailDialog(true);
        }
      })
      .catch((err) => {
        setShowFailDialog(true);
      });
  };

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
      <form
        className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center py-5"
        onSubmit={handleSubmit(onReset)}
      >
        <div className="self-center">
          <div className="text-xl font-medium text-center">RESET PASSWORD</div>
          <div className="self-center text-gray-500 text-center">
            Do not worries, one minute to reset it!
          </div>
        </div>
        <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
          RESET
        </button>
      </form>
      <SimpleDialog
        isOpen={showSuccessDialog}
        setIsOpen={setShowSuccessDialog}
        title="Reset Email Sent"
        content="Please verify in your email inbox for resetting your account password."
        buttonText="Close"
        buttonAction={() => setShowSuccessDialog(false)}
        buttonStatus="success"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Reset Email Request Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default UserPassword;
