import React, { useEffect } from "react";
import notfound from "../notfound.svg";
function NotFound() {
  function Back() {
    window.history.back();
  }
  return (
    <>
      <div className="row">
        <div className="col-1"></div>
        <div className="col">
          <h1 className="notFound">404 Not Found</h1>
          <img
            src={notfound}
            className="img-fluid"
            id="notfound-img"
            alt="Responsive image"
          ></img>
          <br></br>
          <button
            id="back404-btn"
            type="button"
            class="btn btn-triton-secondary btn-lg"
            onClick={Back}
          >
            Back
          </button>
        </div>
        <div className="col-1"></div>
      </div>
    </>
  );
}

export default NotFound;
