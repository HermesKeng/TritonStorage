import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import Axios from "axios";
function RegisterForm(props) {
  async function handleSubmit(e) {
    e.preventDefault();
    if (password != checkPwd) {
      console.log("password should be the same ");
      // jump out alert
      return;
    }
    try {
      const response = await Axios.post("http://localhost:8080/newuser", {
        Username: username,
        Email: email,
        Password: password,
      });
      if (response.data.IsSuccess) {
        localStorage.setItem("tritonStorageToken", response.data.Token);
        localStorage.setItem("tritonStorageUsername", response.data.Username);
        props.setLoggedIn(true);
        console.log("User was successfully created");
      }
    } catch (e) {
      console.log("error");
    }
  }
  const [username, setUsername] = useState();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [checkPwd, setCheckPwd] = useState("");
  const [isPwdSame, setPwdSame] = useState(true);
  useEffect(() => {
    console.log(email);
  }, [email]);
  useEffect(() => {
    if (checkPwd != password && checkPwd.length != 0) {
      setPwdSame(false);
    } else {
      setPwdSame(true);
    }
  }, [password, checkPwd]);
  return (
    <>
      <h3>Register Account</h3>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          Username
          <input
            onChange={e => setUsername(e.target.value)}
            type="text"
            className="form-control"
            placeholder="Enter Username"
          ></input>
        </div>
        <div className="form-group">
          Email address
          <input
            onChange={e => setEmail(e.target.value)}
            type="email"
            className="form-control"
            aria-describedby="emailHelp"
            placeholder="Enter email"
          ></input>
          <small id="emailHelp" className="form-text text-muted">
            We'll never share your email with anyone else.
          </small>
        </div>
        <div className="form-group">
          <label>Password</label>
          <input
            onChange={e => setPassword(e.target.value)}
            type="password"
            className="form-control"
            placeholder="Password"
          ></input>
        </div>
        <div className="form-group">
          <label>Confirm Your Password</label>
          <input
            onChange={e => setCheckPwd(e.target.value)}
            type="password"
            className="form-control"
            placeholder="Password"
          ></input>
        </div>
        {!isPwdSame ? (
          <div className="alert alert-danger" role="alert">
            A simple danger alert—check it out!
          </div>
        ) : (
          ""
        )}
        <button type="submit" className="btn btn-primary">
          Submit
        </button>
        <Link to="#" onClick={() => props.setIsRegister(false)}>
          Login
        </Link>
      </form>
    </>
  );
}

export default RegisterForm;
