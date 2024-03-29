import React, { useEffect, useState } from "react";
import { useHistory } from "react-router-dom";
import useAuthStore from "../store/authStore";
import useIdentityStore from "../store/identityStore";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";
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
      history.push("/manage/users");
      return;
    }
    var errMsg = "";
    switch (res.data?.ErrorMessage) {
      case "invalid.credential": {
        errMsg = "Invalid credential.";
        break;
      }
      case "user.not.found": {
        errMsg = "User not found.";
        break;
      }
      case "invalid.account.status": {
        errMsg = "Invalid account status.";
        break;
      }
      case "invalid.user.role": {
        errMsg = "Invalid user role.";
        break;
      }
      default: {
        errMsg = "Something went wrong, please try again later.";
      }
    }
    setErrMsg(errMsg);
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
    <div className="w-full min-h-screen flex justify-center content-center">
      <div className="w-1/2 hidden lg:flex bg-gradient-to-r from-purple-800 to-indigo-900"></div>
      <div className="w-full lg:w-1/2 bg-gray-700 flex justify-center items-center">
        <form
          onSubmit={handleSubmit(onLogin)}
          className="text-white flex flex-col space-y-5 justify-center self-center"
        >
          <div className="self-center">
            <div className="text-xl font-medium text-center">
              STRAPER ADMIN PORTAL
            </div>
            <div className="self-center text-gray-500 text-center">
              Please enter your credential
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
            <ErrorMessage
              errors={errors}
              name="username"
              as="p"
              className="text-red-600"
            />
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
            <ErrorMessage
              errors={errors}
              name="password"
              as="p"
              className="text-red-600"
            />
          </div>

          {errMsg !== "" && (
            <div className="text-red-600 self-center">{errMsg}</div>
          )}
          <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
            LOG IN
          </button>
        </form>
      </div>
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
