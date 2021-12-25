import React from "react";
import { Route, Redirect } from "react-router-dom";
import useAuthStore from "../../store/authStore";
import useIdentifyStore from "../../store/identityStore";

const NoAuthGuard = ({ component: Component, ...rest }) => {
  const identity = useIdentifyStore((state) => state.identity);

  return (
    <Route
      {...rest}
      render={(props) =>
        identity?.role == "USER" ? (
          <Redirect to="/channels" />
        ) : identity?.role == "ADMIN" ? (
          <Redirect to="/dashboard" />
        ) : (
          <Component {...props} />
        )
      }
    />
  );
};

export default NoAuthGuard;
