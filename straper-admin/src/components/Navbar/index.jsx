import React from "react";
import { FcManager } from "react-icons/fc";
import classNames from "classnames";

import styles from "./Navbar.module.css";
import { Link } from "react-router-dom";
import { logOut } from "../../service/logout";
import useAuthStore from "../../store/authStore";

const Navbar = ({ navigationData, currentRoute, setCurrentRoute }) => {
  const accessToken = useAuthStore((state) => state.accessToken);

  return accessToken && accessToken !== "" ? (
    <nav className={styles.navbar}>
      <span className={styles.logo}>
        <FcManager />
      </span>
      <ul className={styles.navItems}>
        {navigationData.map((item, index) => (
          <Link
            to={"/manage/" + item.toLowerCase()}
            className={classNames([
              styles.navItem,
              currentRoute === item && styles.selectedNavItem,
            ])}
            key={index}
            onClick={() => setCurrentRoute(item)}
          >
            {item}
          </Link>
        ))}
      </ul>
      <button className={styles.actions} onClick={() => logOut(false)}>
        Leave
      </button>
    </nav>
  ) : (
    <></>
  );
};

export default Navbar;
