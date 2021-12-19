import "./App.scss";
import {
  Route,
  BrowserRouter as Router,
  Switch,
  Redirect,
} from "react-router-dom";
import Channel from "./page/Workspace";
import Login from "./page/Login";
import Register from "./page/Register";
import EmailVerify from "./page/EmailVerify";
import UserAuthGuard from "./utils/guards/UserAuthGuard";
import NoAuthGuard from "./utils/guards/NoAuthGuard";
import ResetPassword from "./page/ResetPassword";

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <UserAuthGuard
            path="/channels/:workspaceId/:channelId"
            component={Channel}
          />
          <UserAuthGuard path="/channels" component={Channel} />
          <NoAuthGuard path="/login" component={Login} />
          <NoAuthGuard path="/register" component={Register} />
          <NoAuthGuard path="/reset-password" component={ResetPassword} />
          <Route path="/account/opening/verify">
            <EmailVerify />
          </Route>
          <Route path="/">
            <Redirect to="/login" />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
