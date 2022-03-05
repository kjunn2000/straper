import { ErrorMessage } from "@hookform/error-message";
import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useParams } from "react-router-dom/cjs/react-router-dom.min";
import api from "../axios/api";
import SimpleDialog from "../shared/dialog/SimpleDialog";

const EditUser = () => {
  const {
    handleSubmit,
    register,
    formState: { errors },
    reset,
    watch,
  } = useForm();

  const { userId } = useParams();
  const [user, setUser] = useState();
  const [showSuccessDialog, setShowSuccessDialog] = useState(false);
  const [showFailDialog, setShowFailDialog] = useState(false);
  const [dialogErrMsg, setDialogErrMsg] = useState("");
  const [isPasswdUpdate, setIsPasswdUpdate] = useState(false);
  const watchIsPasswdUpdate = watch("is_passwd_update");

  useEffect(() => {
    fetchUser();
  }, [userId]);

  useEffect(() => resetForm(), [user]);

  const fetchUser = async () => {
    const res = await api.get(`/protected/user/read/${userId}`);
    if (res.data.Success) {
      setUser(res.data.Data);
    }
  };

  const isFormDirty = (data) => {
    return !(
      data.username === user.username &&
      data.email === user.email &&
      data.phone_no === user.phone_no &&
      data.status === user.status &&
      !data.is_passwd_update
    );
  };

  const onUpdate = (data) => {
    if (!isFormDirty(data)) {
      return;
    }
    data.user_id = user.user_id;
    api
      .post("/protected/user/update", data)
      .then((res) => {
        if (res.data.Success) {
          setUser((prev) => ({
            ...prev,
            username: data.username,
            email: data.email,
            phone_no: data.phone_no,
            status: data.status,
          }));
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

  const resetForm = () => {
    if (!user) {
      return;
    }
    setIsPasswdUpdate(false);
    reset({
      username: user.username,
      email: user.email,
      phone_no: user.phone_no,
      status: user.status,
      is_passwd_update: false,
      password: "",
    });
  };

  return (
    <div className="flex justify-center p-5">
      {user && (
        <>
          <form
            className="h-auto rounded-lg flex flex-col space-y-5 bg-white p-5"
            onSubmit={handleSubmit(onUpdate)}
          >
            <div>
              <div className="text-xl font-medium text-center">
                User Information
              </div>
            </div>
            <div>
              <div>Username</div>
              <input
                className="bg-gray-100 p-2 rounded-lg"
                defaultValue={user.username}
                {...register("username", {
                  required: "Username is required.",
                  minLength: {
                    value: 4,
                    message: "Username at least 4 digits.",
                  },
                })}
              />
              <ErrorMessage
                errors={errors}
                name="username"
                as="p"
                className="text-red-500"
              />
            </div>
            <div>
              <div>Email</div>
              <input
                className="bg-gray-100 p-2 rounded-lg"
                defaultValue={user.email}
                {...register("email", {
                  required: "Email is required.",
                  pattern: {
                    value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+.[A-Z]{2,4}$/i,
                    message: "Invalid email format.",
                  },
                })}
              />
              <ErrorMessage
                errors={errors}
                name="email"
                as="p"
                className="text-red-500"
              />
            </div>
            <div>
              <div>Phone No</div>
              <input
                className="bg-gray-100 p-2 rounded-lg"
                type="number"
                defaultValue={user.phone_no}
                {...register("phone_no", {
                  required: "Phone number is required.",
                  pattern: {
                    value: /^[0-9]{10,11}$/,
                    message: "Invalid phone number.",
                  },
                })}
              />
              <ErrorMessage
                errors={errors}
                name="phone_no"
                as="p"
                className="text-red-500"
              />
            </div>
            <div>
              <div>Status</div>
              <select
                {...register("status", {
                  required: "Status is required.",
                })}
                className="p-2 rounded bg-gray-100 hover:cursor-pointer focus:outline-none"
                defaultValue={user.status}
              >
                <option value="ACTIVE">ACTIVE</option>
                <option value="VERIFYING">VERIFYING</option>
                <option value="INACTIVE">INACTIVE</option>
              </select>
              <ErrorMessage
                errors={errors}
                name="status"
                as="p"
                className="text-red-500"
              />
            </div>
            <div className="flex space-x-2 items-center">
              <input
                className="p-2 rounded-lg hover:cursor-pointer"
                {...register("is_passwd_update", {})}
                type="checkbox"
                onChange={(e) => setIsPasswdUpdate(e.target.checked)}
              />
              <span>Update New Password</span>
            </div>
            {isPasswdUpdate && (
              <div>
                <div>New Password</div>
                <input
                  className="bg-gray-100 p-2 rounded-lg"
                  {...register("password", {})}
                  placeholder="Enter new password..."
                  type="password"
                  defaultValue=""
                />
                <ErrorMessage
                  errors={errors}
                  name="password"
                  as="p"
                  className="text-red-500"
                />
              </div>
            )}
            <button
              type="submit"
              className="bg-indigo-400 hover:bg-indigo-200 self-center w-48 p-1 rounded text-white"
            >
              CONFIRM UPDATE
            </button>
            <span
              className="text-sm text-blue-600 hover:text-blue-300 hover:cursor-pointer text-underline self-end"
              onClick={() => resetForm()}
            >
              RESET
            </span>
          </form>
          <SimpleDialog
            isOpen={showSuccessDialog}
            setIsOpen={setShowSuccessDialog}
            title="Update Successfully"
            content="User's lastest information is saved to database."
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
        </>
      )}
    </div>
  );
};

export default EditUser;
