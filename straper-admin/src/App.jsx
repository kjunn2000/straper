import "./App.css";
import ManageUser from "./page/ManageUser";
import ManageWorkspace from "./page/ManageWorkspace";
import Login from "./page/Login";
import {
  Route,
  BrowserRouter as Router,
  Switch,
  Redirect,
} from "react-router-dom";
import AdminAuthGuard from "./utils/guards/AdminAuthGuard";
import NoAuthGuard from "./utils/guards/NoAuthGuard";
import Nav from "./components/Nav/Nav";
import EditUser from "./page/EditUser";
import EditWorkspace from "./page/EditWorkspace";

function App() {
  return (
    <div className="App" className="bg-gray-200 min-h-screen">
      <Router>
        <Nav />
        <Switch>
          <AdminAuthGuard path="/manage/user/:userId" component={EditUser} />
          <AdminAuthGuard path="/manage/users" component={ManageUser} />
          <AdminAuthGuard
            path="/manage/workspace/:workspaceId"
            component={EditWorkspace}
          />
          <AdminAuthGuard
            path="/manage/workspaces"
            component={ManageWorkspace}
          />
          <NoAuthGuard path="/login/timeout" component={Login} />
          <NoAuthGuard path="/login" component={Login} />
          <Route path="/">
            <Redirect to="/login" />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
