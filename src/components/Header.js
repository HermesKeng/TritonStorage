import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import logo from "../logo.svg";
function Header() {
  return (
    <nav className="navbar">
      <Link className="navbar-brand" to="/">
        <img
          src={logo}
          width="30"
          height="30"
          className="d-inline-block align-top"
          id="logo"
          alt=""
        ></img>
        Triton Storage
      </Link>
    </nav>
  );
}

export default Header;
