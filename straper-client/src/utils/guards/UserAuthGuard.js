import React from "react";
import { Route, Redirect } from "react-router-dom";
import useIdentifyStore from "../../store/identityStore";

const UserAuthGuard = ({ component: Component, ...rest }) => {
  const identity = useIdentifyStore((state) => state.identity);
  return (
    <Route
      {...rest}
      render={(props) =>
        identity?.role == "USER" ? (
          <Component {...props} />
        ) : identity?.role == "ADMIN" ? (
          <Redirect to="/dashboard" />
        ) : (
          <Redirect to="/login" />
        )
      }
    />
  );
};

export default UserAuthGuard;
