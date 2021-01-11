import React, { useState } from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Switch, Route } from "react-router-dom";
import "./App.css";
import "./main.scss";
import Header from "./components/Header";
import HomeGuest from "./components/HomeGuest";
import Home from "./components/Home";
import UploadForm from "./components/UploadForm";
import NotFound from "./components/NotFound";
function App() {
  const [loggedIn, setLoggedIn] = useState(
    Boolean(localStorage.getItem("tritonStorageToken"))
  );
  return (
    <BrowserRouter>
      <div className="App">
        <Header></Header>
        <div className="container-fluid">
          <Switch>
            <Route path="/" exact>
              {loggedIn ? (
                <Home
                  name={localStorage.getItem("tritonStorageUsername")}
                  setLoggedIn={setLoggedIn}
                />
              ) : (
                <HomeGuest setLoggedIn={setLoggedIn} />
              )}
            </Route>
            <Route path="/user" exact>
              <Home setLoggedIn={setLoggedIn} />
            </Route>
            <Route path="/newfile" exact>
              <UploadForm />
            </Route>
            <Route>
              <NotFound />
            </Route>
          </Switch>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
