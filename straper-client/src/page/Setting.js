import { useRef, useState } from "react";
import { Tab } from "@headlessui/react";
import { darkGrayBg } from "../utils/style/color";
import { FaWindowClose } from "react-icons/fa";
import AccountInfo from "../components/settings/AccountInfo";
import UserPassword from "../components/settings/UserPassword";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import { logOut } from "../service/logout";
import ActionDialog from "../shared/dialog/ActionDialog";
import { FiSettings } from "react-icons/fi";

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

  const settingSideBar = useRef();

  const handleLogOut = () => {
    setConfirmLogoutDialogOpen(true);
  };

  const toggleSettingSideBar = () => {
    const display = settingSideBar.current.style.display === "" ? "flex" : "";
    settingSideBar.current.style.display = display;
  };

  return (
    <div className="w-full min-h-screen" style={darkGrayBg}>
      <Tab.Group>
        <div className="flex flex-col lg:flex-row lg:flex w-full h-full">
          {/* Mobile View */}
          <div
            className="absolute w-full top-0 text-gray-100 flex justify-between lg:hidden"
            style={{ background: "rgb(32,34,37)" }}
          >
            <a className="block p-4 text-white font-bold skew-x-3 skew-y-3">
              STRAPER
            </a>
            <div>
              <button
                className="mobile-menu-button p-4 focus:outline-none focus:bg-gray-700 hover:bg-indigo-600 transition duration-150"
                onClick={() => toggleSettingSideBar()}
              >
                <FiSettings size={20} />
              </button>
            </div>
          </div>
          <div
            className="w-4/5 lg:w-2/5 h-full flex justify-end absolute lg:relative hidden lg:flex inset-y-0 left-0 z-10"
            style={{ paddingTop: "100px", ...darkGrayBg }}
            ref={settingSideBar}
          >
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
                LOG OUT
              </button>
            </Tab.List>
          </div>
          <Tab.Panels className="w-full lg:w-3/5 h-full">
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
