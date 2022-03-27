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

  const onUpdate = (data) => {
    if (
      data.username === identity.username &&
      data.email === identity.email &&
      data.phone_no === identity.phone_no
    ) {
      return;
    }
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
        <button type="submit" className="bg-indigo-400 self-center w-48 p-1">
          CONFIRM UPDATE
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
