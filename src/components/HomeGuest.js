import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import GoogleLogin from "react-google-login";
import logoImage from "../logo.svg";
import LoginForm from "./LoginForm";
import RegisterForm from "./RegisterForm";
const responseGoogle = response => {
  console.log("success");
  console.log(response);
};
function HomeGuest() {
  const [isRegister, setIsRegister] = useState();
  return (
    <>
      <div className="row">
        <div className="col"></div>
        <div className="col-6 login">
          <img
            src={logoImage}
            className="img-fluid"
            id="main-img"
            alt="Responsive image"
          ></img>
          <h1 className="title">Triton Storage</h1>

          <div className="row login-col">
            <div className="col"></div>
            <GoogleLogin
              clientId="126985904969-a702htpkj5puk0e9oljlr6khepd0brap.apps.googleusercontent.com"
              buttonText="Login"
              onSuccess={responseGoogle}
              onFailure={responseGoogle}
            />
            <div className="col"></div>
          </div>

          <div className="row">
            <div className="col-3"></div>
            <div className="col-6">
              {!isRegister ? (
                <LoginForm setIsRegister={setIsRegister} />
              ) : (
                <RegisterForm setIsRegister={setIsRegister} />
              )}
            </div>
            <div className="col-3"></div>
          </div>
        </div>
        <div className="col"></div>
      </div>
    </>
  );
}

export default HomeGuest;
