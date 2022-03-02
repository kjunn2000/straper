import React, { useCallback } from "react";
import classNames from "classnames";
import { FiUsers } from "react-icons/fi";
import { MdWork } from "react-icons/md";
import { RiLogoutBoxLine } from "react-icons/ri";

import styles from "./Tabbar.module.css";
import { Link } from "react-router-dom";
import useAuthStore from "../../store/authStore";
import { logOut } from "../../service/logout";

const Tabbar = ({ navigationData, currentRoute, setCurrentRoute }) => {
  const accessToken = useAuthStore((state) => state.accessToken);

  const getTabIcon = useCallback((item) => {
    switch (item) {
      case "Users":
        return <FiUsers />;
      case "Workspaces":
        return <MdWork />;
      case "LogOut":
        return <RiLogoutBoxLine />;
    }
  }, []);

  return accessToken && accessToken !== "" ? (
    <nav className={styles.tabbar}>
      {navigationData.map((item, index) => (
        <Link
          to={"/manage/" + item.toLowerCase()}
          key={index}
          className={classNames([
            styles.tabItem,
            currentRoute === item && styles.tabItemActive,
          ])}
          onClick={() => setCurrentRoute(item)}
        >
          <span className={styles.icon}>{getTabIcon(item)}</span>
        </Link>
      ))}
      <span
        className={classNames([styles.tabItem])}
        onClick={() => logOut(false)}
      >
        <span className={styles.icon}>{getTabIcon("LogOut")}</span>
      </span>
    </nav>
  ) : (
    <></>
  );
};

export default Tabbar;
