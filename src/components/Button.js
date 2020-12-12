import React, { useEffect } from "react";

function Button(props) {
  return (
    <button
      className={
        "mdc-button mdc-top-app-bar__action-item mdc-button--outlined " +
        props.name
      }
    >
      <div class="mdc-button__ripple"></div>
      <span className="mdc-button__label">{props.label}</span>
    </button>
  );
}

export default Button;
