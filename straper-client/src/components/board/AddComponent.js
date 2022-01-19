import React, { useRef, useState } from "react";
import { AiOutlineClose } from "react-icons/ai";

const AddComponent = ({ action, type, text }) => {
  const [addMode, setAddMode] = useState(false);
  const inputRef = useRef();

  return (
    <div className="flex flex-col bg-gray-600 bg-opacity-25 rounded-md p-3 m-3 ">
      {addMode ? (
        <div>
          <input className="rounded p-1 mb-3" ref={inputRef} />
          <div className="flex justify-between text-gray-400 hover:text-white cursor-pointer text-sm">
            <button
              onClick={() => {
                action(inputRef.current.value);
                setAddMode(false);
              }}
            >
              Add {type}{" "}
            </button>
            <AiOutlineClose size={20} onClick={() => setAddMode(false)} />
          </div>
        </div>
      ) : (
        <div
          onClick={() => setAddMode(true)}
          className="text-gray-400 hover:text-white cursor-pointer text-sm"
        >
          {text}
        </div>
      )}
    </div>
  );
};

export default AddComponent;
