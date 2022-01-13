import { ErrorMessage } from "@hookform/error-message";
import axios from "axios";
import React, { useRef, useState } from "react";
import { useForm } from "react-hook-form";
import PasswordStrengthBar from "react-password-strength-bar";
import { Link, useHistory } from "react-router-dom/cjs/react-router-dom.min";
import SimpleDialog from "../dialog/SimpleDialog";

const UserPassword = () => {
  const history = useHistory();
  const {
    handleSubmit,
    register,
    watch,
    formState: { errors },
  } = useForm();
  const passwordStrength = useRef();
  const watchPassword = watch("password");
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");

  const onRegister = (data) => {
    axios
      .post("http://localhost:8080/api/v1/account/create", data)
      .then((res) => {
        if (res.data.Success) {
          setShowSuccessDialog(true);
        } else {
          switch (res.data.ErrorMessage) {
            case "username.registered": {
              setDialogErrMsg(
                "Username is registered. Please try other username."
              );
              break;
            }
            case "email.registered": {
              setDialogErrMsg("Email is registered. Please try other email.");
              break;
            }
            case "phone.no.registered": {
              setDialogErrMsg(
                "Phone number is registered. Please try other phone number."
              );
              break;
            }
            case "invalid.username.format": {
              setDialogErrMsg(
                "Phone number format incorrect. Please try again."
              );
              break;
            }
            case "invalid.email.format": {
              setDialogErrMsg("Email format incorrect. Please try again.");
              break;
            }
            case "invalid.phone.no.format": {
              setDialogErrMsg(
                "Phone number format incorrect. Please try again."
              );
              break;
            }
            case "password.too.weak": {
              setDialogErrMsg("Password too weak. Please try again.");
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

  const isPasswordValid = () => {
    return passwordStrength.current.state.score >= 3;
  };

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
      <form
        className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center py-5"
        onSubmit={handleSubmit(onRegister)}
      >
        <div className="self-center">
          <div className="text-xl font-medium text-center">Reset Password</div>
          <div className="self-center text-gray-500 text-center">
            Please note that the old password will be completely removed.
          </div>
        </div>
        <div className="self-center">
          <div>Password</div>
          <input
            type="password"
            className="bg-gray-800 p-2 rounded-lg"
            {...register("password", {
              required: true,
              validate: () => isPasswordValid(),
            })}
          />
          <PasswordStrengthBar
            ref={passwordStrength}
            password={watchPassword}
            className="pt-3"
          />
          {errors?.password && (
            <div className="text-red-500">Password is too weak</div>
          )}
        </div>
        <button
          type="submit"
          className="bg-indigo-400 self-center w-48 p-1 rounded"
        >
          CONFIRM RESET
        </button>
      </form>
      <SimpleDialog
        isOpen={showSuccessDialog}
        setIsOpen={setShowSuccessDialog}
        title="Reset Successfully"
        content="Please log in again with you new password."
        buttonText="Close"
        buttonAction={() => history.push("/login")}
        buttonStatus="success"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Reset Password Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default UserPassword;
