import React, { useState } from "react";
import { Link } from "react-router-dom";
import Axios from "axios";
function ComponentName() {
  const [fileData, setFile] = useState("");
  function handleUpload(e) {
    console.log(e.target.files[0]);
    setFile(e.target.files[0]);
  }

  async function onSubmit() {
    let formdata = new FormData();
    formdata.append("file", fileData);
    console.log(fileData);
    for (var key of formdata.entries()) {
      console.log(key[0] + ", " + key[1].name);
    }
    const response = await Axios.post(
      "http://localhost:8080/newfile",
      formdata
    );
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
            className="form-control"
            type="file"
            id="formFile"
            onChange={handleUpload}
          ></input>
        </div>
        <div className="col-1">
          <button
            type="submit"
            className="btn btn-triton mb-3"
            onClick={onSubmit}
          >
            Upload
          </button>
        </div>
        <div className="col"></div>
      </div>
      <div className="row">
        <div className="col-4">Filename:{fileData.name}</div>
        <div className="col-4">type: {fileData.type}</div>
        <div className="col-4">size: {fileData.size} bytes</div>
      </div>
    </>
  );
}

export default ComponentName;
