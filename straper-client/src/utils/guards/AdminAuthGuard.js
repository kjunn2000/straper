import React from "react";
import { Route, Redirect } from "react-router-dom";
import useIdentityStore from "../../store/identityStore";

const AdminAuthGuard = ({ component: Component, ...rest }) => {
  const identity = useIdentityStore((state) => state.identity);
  return (
    <Route
      {...rest}
      render={(props) =>
        identity?.role === "ADMIN" ? (
          <Component {...props} />
        ) : identity?.role === "USER" ? (
          <Redirect to="/channel" />
        ) : (
          <Redirect to="/login" />
        )
      }
    />
  );
};

export default AdminAuthGuard;
