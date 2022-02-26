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

function App() {
  return (
    <div className="App" className="bg-gray-200 h-screen">
      <Router>
        <Nav />
        <Switch>
          <AdminAuthGuard path="/manage/user" component={ManageUser} />
          <AdminAuthGuard
            path="/manage/workspace"
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
