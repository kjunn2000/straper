import './App.scss';
import { Route, BrowserRouter as Router, Switch, Redirect } from 'react-router-dom';
import Chat from './page/Chat';
import Login from './page/Login';
import Register from './page/Register';

function App() {

  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/workspace">
            <Chat/>
          </Route>
          <Route path="/login">
            <Login/>
          </Route>
          <Route path="/register">
            <Register/>
          </Route>
          <Route path="/">
            <Redirect to="/login"/>
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
