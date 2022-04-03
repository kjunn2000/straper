import React, { useRef, useState } from "react";
import { Link, useHistory } from "react-router-dom";
import PasswordStrengthBar from "react-password-strength-bar";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";
import api from "../axios/api";
import SimpleDialog from "../shared/dialog/SimpleDialog";

const Register = () => {
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
    api
      .post("/account/create", data)
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
              setDialogErrMsg("Invalid username format.");
              break;
            }
            case "invalid.email.format": {
              setDialogErrMsg("Invalid email format.");
              break;
            }
            case "invalid.phone.no.format": {
              setDialogErrMsg("Invalid phone number format.");
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
      .catch(() => {
        setShowFailDialog(true);
      });
  };

  const isPasswordValid = () => {
    return passwordStrength.current.state.score >= 3;
  };

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full min-h-screen flex justify-center content-center">
      <form
        className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-72 md:w-96 h-auto max-h-full justify-center self-center py-5"
        onSubmit={handleSubmit(onRegister)}
      >
        <div className="self-center">
          <div className="text-xl font-medium text-center">WELCOME BACK</div>
          <div className="self-center text-gray-500 text-center">
            Good To See You Again, Friends!
          </div>
        </div>
        <div className="self-center">
          <div>Username</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            {...register("username", {
              required: "Username is required.",
              minLength: {
                value: 4,
                message: "Username at least 4 chars.",
              },
            })}
          />
          <ErrorMessage errors={errors} name="username" as="p" />
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
        <div className="self-center">
          <div>Phone No</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            type="number"
            {...register("phone_no", {
              required: "Phone number is required.",
              pattern: {
                value: /^[0-9]{10,11}$/,
                message: "Invalid phone number.",
              },
            })}
          />
          <ErrorMessage errors={errors} name="phone_no" as="p" />
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
          {errors?.password && errors?.password.type === "required" ? (
            <div className="text-red-500">Password is required.</div>
          ) : errors?.password && errors?.password.type === "validate" ? (
            <div className="text-red-500">Password is too weak.</div>
          ) : (
            <></>
          )}
        </div>
        <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
          REGISTER NOW
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
        title="Register Successfully"
        content="Thank you for registering account in Straper, please verify your 
          email in your email inbox to complete the registration."
        buttonText="Close"
        buttonAction={() => history.push("/login")}
        buttonStatus="success"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Registered Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default Register;
