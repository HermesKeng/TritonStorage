import React, { useState } from "react";
import { Link } from "react-router-dom";

function ComponentName() {
  const [file, setFile] = React.useState("");
  function handleUpload(e) {
    setFile(e.target.files[0]);
  }
  return (
    <>
      <div className="row first">
        <div className="col-4">
          <h1> Upload File</h1>
        </div>
        <div className="col-4"></div>
        <div className="col-4">
          <Link to="/" type="button" className="btn btn-outline-triton">
            Back
          </Link>
        </div>
      </div>
      <div className="row">
        <div className="col-1"></div>
        <div className="col-3">
          <input
            class="form-control"
            type="file"
            id="formFile"
            onChange={handleUpload}
          ></input>
        </div>
        <div className="col-1">
          <button type="submit" className="btn btn-triton mb-3">
            Upload
          </button>
        </div>
        <div className="col"></div>
      </div>
      <div className="row">
        <div className="col-4">Filename:{file.name}</div>
        <div className="col-4">type: {file.type}</div>
        <div className="col-4">size: {file.size} bytes</div>
      </div>
    </>
  );
}

export default ComponentName;
