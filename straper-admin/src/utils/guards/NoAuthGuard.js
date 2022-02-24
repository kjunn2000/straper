import React from "react";
import { Route, Redirect } from "react-router-dom";
import useIdentityStore from "../../store/identityStore";

const NoAuthGuard = ({ component: Component, ...rest }) => {
  const identity = useIdentityStore((state) => state.identity);

  return (
    <Route
      {...rest}
      render={(props) =>
        identity?.role === "ADMIN" ? (
          <Redirect to="/manage/user" />
        ) : (
          <Component {...props} />
        )
      }
    />
  );
};

export default NoAuthGuard;
