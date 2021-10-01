import axios from "../axios/api";
import React, { useRef, useState } from "react";
import { Link, useHistory } from "react-router-dom";
import PasswordStrengthBar from "react-password-strength-bar";
import { useForm } from "react-hook-form";
import SimpleDialog from "../components/dialog/SimpleDialog";

const Register = () => {
  const history = useHistory();
  const {
    handleSubmit,
    register,
    watch,
    getValues,
    formState: { errors },
  } = useForm();
  const passwordStrength = useRef();
  const watchPassword = watch("loginform.password");
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);

  const onRegister = () => {
    axios
      .post(
        "http://localhost:8080/api/v1/account/opening",
        getValues("loginform")
      )
      .then((res) => {
        if (res.data.Success) {
          setShowSuccessDialog(true);
        }
        setShowFailDialog(true);
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
          <div className="text-xl font-medium text-center">WELCOME BACK</div>
          <div className="self-center text-gray-500 text-center">
            Good To See You Again, Friends!
          </div>
        </div>
        <div className="self-center">
          <div>Username</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            {...register("loginform.username", {
              required: true,
              minLength: 4,
            })}
          />
          {errors?.loginform?.username?.type === "required" && (
            <div className="text-red-500">Username cannot be empty</div>
          )}
          {errors?.loginform?.username?.type === "minLength" && (
            <div className="text-red-500">Username at least 4 digits</div>
          )}
        </div>
        <div className="self-center">
          <div>Email</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            {...register("loginform.email", {
              required: true,
              pattern: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+.[A-Z]{2,4}$/i,
            })}
          />
          {errors?.loginform?.email && (
            <div className="text-red-500">Incorrect email format</div>
          )}
        </div>
        <div className="self-center">
          <div>Phone No</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            type="number"
            {...register("loginform.phone_no", {
              required: true,
              pattern: /^[0-9]{10,11}$/,
            })}
          />
          {errors?.loginform?.phone_no && (
            <div className="text-red-500">Invalid phone no</div>
          )}
        </div>
        <div className="self-center">
          <div>Password</div>
          <input
            type="password"
            className="bg-gray-800 p-2 rounded-lg"
            {...register("loginform.password", {
              required: true,
              validate: () => isPasswordValid(),
            })}
          />
          <PasswordStrengthBar
            ref={passwordStrength}
            password={watchPassword}
            className="pt-3"
          />
          {errors?.loginform?.password && (
            <div className="text-red-500">Password is too weak</div>
          )}
        </div>
        <button
          type="submit"
          className="bg-indigo-400 self-center w-48 p-1"
          onClick={() => console.log(errors)}
        >
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
        setIsOpen={setShowFailDialog}
        title="Register Successfully"
        content="Thank you for create an account in Straper, hope you enjoy all the features that we provided."
        buttonText="Close"
        buttonAction={() => history.push("/login")}
        buttonStatus="success"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Registered Fail"
        content="Thank you for create an account in Straper, please try again later."
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default Register;
