import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Switch, Route } from "react-router-dom";
import "./App.css";
import "./main.scss";
import Header from "./components/Header";
import HomeGuest from "./components/HomeGuest";
import Home from "./components/Home";
function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Header></Header>
        <div className="container-fluid">
          <Switch>
            <Route path="/" exact>
              <HomeGuest />
            </Route>
            <Route path="/user" exact>
              <Home />
            </Route>
          </Switch>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
