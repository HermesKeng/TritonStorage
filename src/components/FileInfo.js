import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import Axios from "axios";
function FileInfo(props) {
  function download(file) {
    const url = window.URL.createObjectURL(new Blob([file]));
    var element = document.createElement("a");
    element.href = url;
    element.setAttribute("download", `${props.filename}`);
    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
  }
  const downloadFile = function (e) {
    var link = `/${localStorage.getItem("tritonStorageUsername")}/files/${
      e.target.id
    }`;
    async function fetchPostData() {
      try {
        const response = await Axios.get(link, {
          responseType: "blob",
        });
        download(response.data);
      } catch (error) {
        console.log(error);
      }
    }
    fetchPostData();
  };

  const deleteFile = function (e) {
    alert("delete file :" + e.target.id);
    var link = `/${localStorage.getItem("tritonStorageUsername")}/files/${
      e.target.id
    }`;
    async function deleteData() {
      try {
        const response = await Axios.delete(link);
      } catch (error) {}
    }
    deleteData();
    props.delItem(e.target.id);
  };
  return (
    <>
      <tr>
        <td scope="row">{props.idx + 1}</td>
        <td>{props.filename}</td>
        <td>{props.type}</td>
        <td>{props.version}</td>
        <td>
          <div className="clickArea">
            <Link
              id={props.id}
              onClick={downloadFile}
              className="btn btn-outline-triton-secondary"
            >
              <svg
                id={props.id}
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-cloud-download"
                viewBox="0 0 16 16"
              >
                <path
                  id={props.id}
                  fill-rule="evenodd"
                  d="M4.406 1.342A5.53 5.53 0 0 1 8 0c2.69 0 4.923 2 5.166 4.579C14.758 4.804 16 6.137 16 7.773 16 9.569 14.502 11 12.687 11H10a.5.5 0 0 1 0-1h2.688C13.979 10 15 8.988 15 7.773c0-1.216-1.02-2.228-2.313-2.228h-.5v-.5C12.188 2.825 10.328 1 8 1a4.53 4.53 0 0 0-2.941 1.1c-.757.652-1.153 1.438-1.153 2.055v.448l-.445.049C2.064 4.805 1 5.952 1 7.318 1 8.785 2.23 10 3.781 10H6a.5.5 0 0 1 0 1H3.781C1.708 11 0 9.366 0 7.318c0-1.763 1.266-3.223 2.942-3.593.143-.863.698-1.723 1.464-2.383z"
                />
                <path
                  id={props.id}
                  fill-rule="evenodd"
                  d="M7.646 15.854a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708L8.5 14.293V5.5a.5.5 0 0 0-1 0v8.793l-2.146-2.147a.5.5 0 0 0-.708.708l3 3z"
                />
              </svg>
            </Link>
          </div>
        </td>
        <td>
          <div className="clickArea">
            <Link
              id={props.id}
              onClick={deleteFile}
              className="btn btn-outline-triton-secondary"
            >
              <svg
                id={props.id}
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-trash2-fill"
                viewBox="0 0 16 16"
              >
                <path
                  id={props.id}
                  d="M2.037 3.225A.703.703 0 0 1 2 3c0-1.105 2.686-2 6-2s6 .895 6 2a.702.702 0 0 1-.037.225l-1.684 10.104A2 2 0 0 1 10.305 15H5.694a2 2 0 0 1-1.973-1.671L2.037 3.225zm9.89-.69C10.966 2.214 9.578 2 8 2c-1.58 0-2.968.215-3.926.534-.477.16-.795.327-.975.466.18.14.498.307.975.466C5.032 3.786 6.42 4 8 4s2.967-.215 3.926-.534c.477-.16.795-.327.975-.466-.18-.14-.498-.307-.975-.466z"
                />
              </svg>
            </Link>
          </div>
        </td>
      </tr>
    </>
  );
}

export default FileInfo;
