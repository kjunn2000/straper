import "./App.scss";
import {
  Route,
  BrowserRouter as Router,
  Switch,
  Redirect,
} from "react-router-dom";
import Login from "./page/Login";
import Register from "./page/Register";
import EmailVerify from "./page/EmailVerify";
import UserAuthGuard from "./utils/guards/UserAuthGuard";
import NoAuthGuard from "./utils/guards/NoAuthGuard";
import ResetPasswordRequest from "./page/ResetPasswordRequest";
import ResetPassword from "./page/ResetPassword";
import Workspace from "./page/Workspace";
import Setting from "./page/Setting";
import TaskBoard from "./page/TaskBoard";
import LogOutNotice from "./page/LogOutNotice";
import Bug from "./page/Bug";
import IssueDetail from "./page/IssueDetail";

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <UserAuthGuard
            path="/channel/:workspaceId/:channelId"
            component={Workspace}
          />
          <UserAuthGuard path="/channel" component={Workspace} />
          <UserAuthGuard path="/setting" component={Setting} />
          <UserAuthGuard path="/board" component={TaskBoard} />
          <UserAuthGuard path="/bug" component={Bug} />
          <UserAuthGuard path="/issue/:issueId" component={IssueDetail} />
          <UserAuthGuard path="/logout-notice" component={LogOutNotice} />
          <NoAuthGuard path="/login/timeout" component={Login} />
          <NoAuthGuard path="/login" component={Login} />
          <NoAuthGuard path="/register" component={Register} />
          <NoAuthGuard
            path="/reset-password"
            component={ResetPasswordRequest}
          />
          <Route path="/account/opening/verify">
            <EmailVerify />
          </Route>
          <Route path="/account/password/update">
            <ResetPassword />
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
