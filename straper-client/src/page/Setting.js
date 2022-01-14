import { useState } from "react";
import { Tab } from "@headlessui/react";
import { darkGrayBg } from "../utils/style/color";
import { FaWindowClose } from "react-icons/fa";
import AccountInfo from "../components/settings/AccountInfo";
import UserPassword from "../components/settings/UserPassword";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import { logOut } from "../service/logout";
import ActionDialog from "../components/dialog/ActionDialog";
import { set } from "react-hook-form";

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

function Setting() {
  let [categories] = useState({
    "Account Info": <AccountInfo />,
    "User Password": <UserPassword />,
  });

  const history = useHistory();
  const [isConfirmLogoutDialogOpen, setConfirmLogoutDialogOpen] =
    useState(false);

  const handleLogOut = () => {
    setConfirmLogoutDialogOpen(true);
  };

  return (
    <div className="w-full h-screen" style={darkGrayBg}>
      <Tab.Group vertical={true}>
        <div className="grid grid-cols-3 w-full h-full">
          <div className="flex justify-end" style={{ paddingTop: "100px" }}>
            <FaWindowClose
              size="40"
              className="text-indigo-500 mr-5 cursor-pointer"
              onClick={() => history.push("/channel")}
            />
            <Tab.List className="flex flex-col w-1/3 h-full">
              <div className="text-white font-bold text-sm">USER SETTINGS</div>
              {Object.keys(categories).map((category) => (
                <Tab
                  key={category}
                  className={({ selected }) =>
                    classNames(
                      "w-full py-3 px-5 text-sm font-medium text-left text-gray-300 rounded transition-all duration-300",
                      selected
                        ? "bg-gray-600 drop-shadow-lg"
                        : "hover:bg-gray-700"
                    )
                  }
                >
                  {category.toUpperCase()}
                </Tab>
              ))}
              <button
                onClick={() => handleLogOut()}
                className="w-full py-3 px-5 text-sm font-medium text-left text-gray-300 rounded transition-all duration-300 
                    drop-shadow-lg hover:bg-red-600"
              >
                Log Out
              </button>
            </Tab.List>
          </div>
          <Tab.Panels className="w-full h-full col-span-2">
            {Object.values(categories).map((category, idx) => (
              <Tab.Panel key={idx} className={classNames("")}>
                {category}
              </Tab.Panel>
            ))}
          </Tab.Panels>
        </div>
      </Tab.Group>
      <ActionDialog
        isOpen={isConfirmLogoutDialogOpen}
        setIsOpen={setConfirmLogoutDialogOpen}
        title="Confirm Logout Straper"
        content="You will be completely log out from straper once you confirm."
        buttonText="Log Out"
        buttonStatus="fail"
        buttonAction={() => logOut()}
        closeButtonText="Close"
      />
    </div>
  );
}

export default Setting;
