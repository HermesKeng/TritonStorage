import React, { useState } from "react";
import FileInfo from "./FileInfo.js";
import { GoogleLogout } from "react-google-login";
const fakeFile = [
  { id: 1, filename: "hello.txt", type: "document", version: 1 },
  { id: 2, filename: "happy.jpg", type: "image", version: 3 },
  { id: 3, filename: "test.mp3", type: "audio", version: 2 },
  { id: 4, filename: "tapie.mp4", type: "video", version: 4 },
];
function Home(props) {
  function LogOut() {
    props.setLoggedIn(false);
    localStorage.removeItem("tritonStorageToken");
    localStorage.removeItem("tritonStorageUsername");
  }
  return (
    <>
      <div className="row first">
        <div className="col-4">
          <h1> Hello! {props.name}</h1>
        </div>
        <div className="col-4"></div>
        <div className="col-4">
          <button
            type="button"
            onClick={LogOut}
            className="btn btn-outline-triton"
          >
            Log Out
          </button>
        </div>
      </div>
      <div className="row">
        <div className="col-1"></div>
        <div className="col">
          <table className="table">
            <thead>
              <tr>
                <th scope="col">ID</th>
                <th scope="col">Filename</th>
                <th scope="col">Type</th>
                <th scope="col">Version</th>
                <th scope="col">Download</th>
              </tr>
            </thead>
            {fakeFile.map(file => (
              <FileInfo
                id={file.id}
                filename={file.filename}
                type={file.type}
                version={file.version}
              />
            ))}
          </table>
        </div>
        <div className="col-1"></div>
      </div>
    </>
  );
}

export default Home;
