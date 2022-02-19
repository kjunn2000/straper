import React from "react";
import { Route, Redirect } from "react-router-dom";
import useIdentityStore from "../../store/identityStore";

const NoAuthGuard = ({ component: Component, ...rest }) => {
  const identity = useIdentityStore((state) => state.identity);

  return (
    <Route
      {...rest}
      render={(props) =>
        identity?.role === "USER" ? (
          <Redirect to="/channel" />
        ) : identity?.role === "ADMIN" ? (
          <Redirect to="/dashboard" />
        ) : (
          <Component {...props} />
        )
      }
    />
  );
};

export default NoAuthGuard;
