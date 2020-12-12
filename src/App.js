import React from "react";
import "./App.css";
import "./main.scss";
import Header from "./components/Header";
import Guest from "./components/Guest";
function App() {
  return (
    <div className="App">
      <Header></Header>
      <div className="container-fluid">
        <Guest></Guest>
      </div>
    </div>
  );
}

export default App;
