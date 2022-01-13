import { useState } from "react";
import { Tab } from "@headlessui/react";
import { darkGrayBg } from "../utils/style/color";
import AccountInfo from "../components/settings/AccountInfo";
import UserPassword from "../components/settings/UserPassword";

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

function Setting() {
  let [categories] = useState({
    "Account Info": <AccountInfo />,
    "User Password": <UserPassword />,
  });

  return (
    <div className="w-full h-screen" style={darkGrayBg}>
      <Tab.Group vertical={true}>
        <div className="grid grid-cols-3 w-full h-full">
          <div className="flex justify-end">
            <Tab.List
              className="flex flex-col w-1/3 h-full"
              style={{ paddingTop: "100px" }}
            >
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
    </div>
  );
}

export default Setting;
