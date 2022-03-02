import React from "react";
import loader from "../assets/img/spinner.gif";

const Loader = () => {
  return (
    <div className="w-full h-auto flex justify-center items-center">
      <img src={loader} alt="Loader" />
    </div>
  );
};

export default Loader;
