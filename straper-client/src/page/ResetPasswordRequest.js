import React, { useState } from "react";
import { Link, useHistory } from "react-router-dom";
import { useForm } from "react-hook-form";
import SimpleDialog from "../components/dialog/SimpleDialog";
import { ErrorMessage } from "@hookform/error-message";
import api from "../axios/api";

const ResetPassword = () => {
  const history = useHistory();
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm();
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");

  const onReset = (data) => {
    api
      .post("/account/reset-password/create", data)
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
        <div className="self-center">
          <div>Email</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            {...register("email", {
              required: "Email is required.",
              pattern: {
                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+.[A-Z]{2,4}$/i,
                message: "Invalid email format.",
              },
            })}
          />
          <ErrorMessage errors={errors} name="email" as="p" />
        </div>
        <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
          RESET
        </button>
        <Link
          to="/login"
          className="text-indigo-300 self-center cursor-pointer hover:text-indigo-500"
        >
          Go to login ?
        </Link>
      </form>
      <SimpleDialog
        isOpen={showSuccessDialog}
        setIsOpen={setShowSuccessDialog}
        title="Reset Email Sent"
        content="Please verify in your email inbox for resetting your account password."
        buttonText="Close"
        buttonAction={() => history.push("/login")}
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

export default ResetPassword;
