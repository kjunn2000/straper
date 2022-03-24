import React, { useRef, useState } from "react";
import { Link, useHistory } from "react-router-dom";
import PasswordStrengthBar from "react-password-strength-bar";
import { useForm } from "react-hook-form";
import api from "../axios/api";
import SimpleDialog from "../shared/dialog/SimpleDialog";

const ResetPassword = () => {
  const history = useHistory();
  const {
    handleSubmit,
    register,
    watch,
    getValues,
    formState: { errors },
  } = useForm();
  const passwordStrength = useRef();
  const watchPassword = watch("password");
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");

  const onReset = (data) => {
    var requestData = {
      token_id: history.location.pathname.split("/").pop(),
      password: data.password,
    };
    api
      .post("/account/password/update", requestData)
      .then((res) => {
        if (res.data.Success) {
          setShowSuccessDialog(true);
        } else {
          switch (res.data.ErrorMessage) {
            case "reset.password.token.not.found": {
              setDialogErrMsg(
                "Invalid reset password token. Please try again."
              );
              break;
            }
            case "reset.password.token.expired": {
              setDialogErrMsg(
                "Reset password token expired. Please send the request again."
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

  const isPasswordMatch = () => {
    return getValues("password") === getValues("confirmedPassword");
  };

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
      <form
        className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center py-5"
        onSubmit={handleSubmit(onReset)}
      >
        <div className="self-center">
          <div className="text-xl font-medium text-center">WELCOME BACK</div>
          <div className="self-center text-gray-500 text-center">
            Please enter your new password for your account.
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
          {errors?.password && (
            <div className="text-red-500">Password is too weak.</div>
          )}
        </div>
        <div className="self-center">
          <div>Confirmed Password</div>
          <input
            type="password"
            className="bg-gray-800 p-2 rounded-lg"
            {...register("confirmedPassword", {
              required: true,
              validate: () => isPasswordMatch(),
            })}
          />
          {errors?.confirmedPassword && (
            <div className="text-red-500">Password not match</div>
          )}
        </div>
        <PasswordStrengthBar
          ref={passwordStrength}
          password={watchPassword}
          className="pt-3"
        />
        <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
          CONFIRM RESET
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
        title="Password Reset Successfully"
        content="Thank you for trusting Straper, you are free to login using your new password."
        buttonText="Close"
        buttonAction={() => history.push("/login")}
        buttonStatus="success"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Passord Reset Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default ResetPassword;
