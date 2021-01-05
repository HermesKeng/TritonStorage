import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Axios from "axios";
function LoginForm(props) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  async function handleSubmit(e) {
    e.preventDefault();
    try {
      const response = await Axios.post("http://localhost:8080/users", {
        Email: email,
        Password: password,
      });
      console.log(response);
      if (response.data.IsSuccess) {
        console.log("User Successuflly Login");
        localStorage.setItem("tritonStorageToken", response.data.Token);
        localStorage.setItem("tritonStorageUsername", response.data.Username);
        props.setLoggedIn(true);
      } else {
        console.log("Please check your account and password !");
      }
    } catch (e) {
      console.log("error");
    }
  }
  return (
    <form onSubmit={handleSubmit}>
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
      <button type="submit" className="btn btn-primary">
        Submit
      </button>
      <Link
        to="#"
        onClick={() => {
          props.setIsRegister(true);
        }}
      >
        Register
      </Link>
    </form>
  );
}

export default LoginForm;
