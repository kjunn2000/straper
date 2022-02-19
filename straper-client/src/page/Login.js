import React, { useEffect, useState } from "react";
import { Link, useHistory } from "react-router-dom";
import useAuthStore from "../store/authStore";
import useIdentityStore from "../store/identityStore";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";
import "./login.scss";
import api from "../axios/api";
import SimpleDialog from "../shared/dialog/SimpleDialog";

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const history = useHistory();

  const setAccessToken = useAuthStore((state) => state.setAccessToken);
  const setIdentity = useIdentityStore((state) => state.setIdentity);

  const [errMsg, setErrMsg] = useState("");
  const [showTimeoutDialog, setShowTimeoutDialog] = useState(false);

  useEffect(() => {
    if (window.location.href.endsWith("/timeout")) {
      setShowTimeoutDialog(true);
    }
  }, []);

  const onLogin = async (data) => {
    const res = await api.post("/auth/login", data, { withCredentials: true });
    if (res.data?.Success) {
      await updateAuthAndIdentityState(
        res.data?.Data.access_token,
        res.data?.Data.user
      );
      history.push("/channel");
    } else if (res.data?.ErrorMessage === "invalid.credential") {
      updateErrMsg("Invalid credenital.");
    } else if (res.data?.ErrorMessage === "user.not.found") {
      updateErrMsg("User not found.");
    } else if (res.data?.ErrorMessage === "invalid.account.status") {
      updateErrMsg("Invalid account status.");
    }
  };

  const updateAuthAndIdentityState = async (accessToken, identity) => {
    setIdentity(identity);
    setAccessToken(accessToken);
  };

  const updateErrMsg = (msg) => {
    setErrMsg(msg);
    setTimeout(() => {
      setErrMsg("");
    }, 5000);
  };

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full h-screen flex justify-center content-center">
      <form
        onSubmit={handleSubmit(onLogin)}
        className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-96 h-auto justify-center self-center py-5"
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
              minLength: { value: 4, message: "Username at leat 4 digits." },
            })}
          />
          <ErrorMessage errors={errors} name="username" as="p" />
        </div>
        <div className="self-center">
          <div>Password</div>
          <input
            type="password"
            className="bg-gray-800 p-2 rounded-lg"
            {...register("password", {
              required: "Password is required.",
            })}
          />
          <ErrorMessage errors={errors} name="password" as="p" />
        </div>

        {errMsg !== "" && (
          <div className="text-red-600 self-center">{errMsg}</div>
        )}
        <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
          LET'S GO
        </button>
        <Link
          to="/register"
          className="text-indigo-300 self-center cursor-pointer hover:text-indigo-500"
        >
          Register an account
        </Link>
        <Link
          to="/reset-password"
          className="text-indigo-300 self-center cursor-pointer hover:text-indigo-500"
        >
          Forget password ?
        </Link>
      </form>
      <SimpleDialog
        isOpen={showTimeoutDialog}
        setIsOpen={setShowTimeoutDialog}
        title="Time Out"
        content="Session is timeout. Please login again."
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default Login;
