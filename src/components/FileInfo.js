import React, { useEffect } from "react";

function FileInfo(props) {
  return (
    <>
      <tr id={props.id}>
        <td scope="row">{props.id}</td>
        <td>{props.filename}</td>
        <td>{props.type}</td>
        <td>{props.version}</td>
      </tr>
    </>
  );
}

export default FileInfo;
