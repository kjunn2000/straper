import { ErrorMessage } from "@hookform/error-message";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import api from "../../axios/api";
import SimpleDialog from "../../shared/dialog/SimpleDialog";
import useIdentityStore from "../../store/identityStore";

const AccountInfo = () => {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm();
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");
  const [isEmailUpdate, setEmailUpdate] = useState(false);
  const identity = useIdentityStore((state) => state.identity);
  const setIdentity = useIdentityStore((state) => state.setIdentity);
  const [isFormSubmit, setIsFormSubmit] = useState(false);

  const onUpdate = (data) => {
    if (
      data.username === identity.username &&
      data.email === identity.email &&
      data.phone_no === identity.phone_no
    ) {
      return;
    }
    setIsFormSubmit(true);
    setEmailUpdate(data.email !== identity.email);
    api
      .post("/protected/account/update", data)
      .then((res) => {
        if (res.data.Success) {
          const newIdentity = {
            ...identity,
            username: data.username,
            email: data.email,
            phone_no: data.phone_no,
          };
          setIdentity(newIdentity);
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
              setDialogErrMsg("Username format incorrect. Please try again.");
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
            default: {
              setDialogErrMsg("Something went wrong. Please try again.");
            }
          }
          setShowFailDialog(true);
        }
        setIsFormSubmit(false);
      })
      .catch((err) => {
        setShowFailDialog(true);
      });
  };

  const successDialogContent =
    "Your account information has updated successful to the system. " +
    (isEmailUpdate
      ? "Please verify your new email at your email inbox before the next login."
      : "");

  return (
    <div className="bg-gradient-to-r from-purple-600 to-gray-900 w-full min-h-screen flex justify-center content-center">
      <form
        className="bg-gray-700 rounded-lg text-white flex flex-col space-y-5 w-72 md:w-96 h-auto justify-center self-center py-5"
        onSubmit={handleSubmit(onUpdate)}
      >
        <div className="self-center">
          <div className="text-xl font-medium text-center">
            Account Information
          </div>
          <div className="self-center text-gray-500 text-center">
            Feel free to update your latest profile.
          </div>
        </div>
        <div className="self-center">
          <div>Username</div>
          <input
            className="bg-gray-800 p-2 rounded-lg"
            defaultValue={identity.username}
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
            defaultValue={identity.email}
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
            defaultValue={identity.phone_no}
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
        <button
          type="submit"
          disabled={isFormSubmit}
          className="bg-indigo-400 self-center w-48 p-1 flex justify-center items-center"
        >
          {!isFormSubmit ? (
            <span> UPDATE </span>
          ) : (
            <>
              <svg
                role="status"
                className="mr-2 w-6 h-6 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
                viewBox="0 0 100 101"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                  fill="currentColor"
                />
                <path
                  d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                  fill="currentFill"
                />
              </svg>
              <span> LOADING... </span>
            </>
          )}
        </button>
      </form>
      <SimpleDialog
        isOpen={showSuccessDialog}
        setIsOpen={setShowSuccessDialog}
        title="Update Successfully"
        content={successDialogContent}
        buttonText="Close"
        buttonAction={() => setShowSuccessDialog(false)}
        buttonStatus="success"
      />

      <SimpleDialog
        isOpen={showFailDialog}
        setIsOpen={setShowFailDialog}
        title="Update Fail"
        content={dialogErrMsg}
        buttonText="Close"
        buttonStatus="fail"
      />
    </div>
  );
};

export default AccountInfo;
