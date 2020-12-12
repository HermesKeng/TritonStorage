import React, { useEffect } from "react";
import logoImage from "../logo.svg";
function Guest() {
  return (
    <>
      <div className="row">
        <div className="col"></div>
        <div className="col-6 login">
          <img
            src={logoImage}
            class="img-fluid"
            id="main-img"
            alt="Responsive image"
          ></img>
          <h1 className="title">Triton Storage</h1>
          <div className="row">
            <div className="col"></div>
            <form>
              <div class="form-group">
                Email address
                <input
                  type="email"
                  class="form-control"
                  aria-describedby="emailHelp"
                  placeholder="Enter email"
                ></input>
                <small id="emailHelp" class="form-text text-muted">
                  We'll never share your email with anyone else.
                </small>
              </div>
              <div class="form-group">
                <label>Password</label>
                <input
                  type="password"
                  class="form-control"
                  placeholder="Password"
                ></input>
              </div>
              <button type="submit" class="btn btn-primary">
                Submit
              </button>
            </form>
            <div className="col"></div>
          </div>
        </div>
        <div className="col"></div>
      </div>
    </>
  );
}

export default Guest;
