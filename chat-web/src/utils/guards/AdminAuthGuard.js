import React from "react";
import { Route, Redirect } from "react-router-dom";
import useIdentifyStore from "../../store/identityStore";

const AdminAuthGuard = ({ component: Component, ...rest }) => {
  const identity = useIdentifyStore((state) => state.identity);
  return (
    <Route
      {...rest}
      render={(props) =>
        identity?.role == "ADMIN" ? (
          <Component {...props} />
        ) : identity?.role == "USER" ? (
          <Redirect to="/channels" />
        ) : (
          <Redirect to="/login" />
        )
      }
    />
  );
};

export default AdminAuthGuard;
