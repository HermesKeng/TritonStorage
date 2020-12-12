import React, { useEffect } from "react";
import Button from "./Button";
import logo from "../logo.svg";
function Header() {
  return (
    <nav className="navbar">
      <a className="navbar-brand" href="#">
        <img
          src={logo}
          width="30"
          height="30"
          className="d-inline-block align-top"
          id="logo"
          alt=""
        ></img>
        Triton Storage
      </a>
    </nav>
  );
}

export default Header;
