import React, { useState, useEffect } from "react";
import FileInfo from "./FileInfo.js";
import { Link } from "react-router-dom";
import { GoogleLogout } from "react-google-login";
import LoadingDotIcon from "./LoadingDotIcon";
import Axios from "axios";
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
  const [posts, setPosts] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  useEffect(() => {
    async function fetchPostData() {
      try {
        const response = await Axios.get(
          `/${localStorage.getItem("tritonStorageUsername")}/files`,
          {
            headers: {
              "x-user": localStorage.getItem("tritonStorageUsername"),
            },
          }
        );
        setPosts(response.data);
        setIsLoading(false);
      } catch (error) {
        console.log("error or the request is cancelled");
      }
    }
    fetchPostData();
  }, []);
  if (isLoading) {
    return <LoadingDotIcon />;
  } else {
    return (
      <>
        <div className="row first">
          <div className="col-4">
            <h1> Hello! {props.name}</h1>
          </div>
          <div className="col-4"></div>
          <div className="col-4">
            <Link
              to="/newfile"
              type="button"
              className="btn btn-outline-triton"
            >
              Upload File
            </Link>
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
              {posts.map((file, index) => (
                <FileInfo
                  id={file.Id}
                  idx={index}
                  key={file.Id}
                  filename={file.Filename}
                  type={file.Type}
                  version={1}
                />
              ))}
            </table>
          </div>
          <div className="col-1"></div>
        </div>
      </>
    );
  }
}

export default Home;
