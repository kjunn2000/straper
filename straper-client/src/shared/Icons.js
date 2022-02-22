import { AiFillBug } from "react-icons/ai";
import { BiTask } from "react-icons/bi";
import {
  BsSubtract,
  BsFillBookmarkFill,
  BsFillLightningFill,
} from "react-icons/bs";

export function SortIcon({ className }) {
  return (
    <svg
      className={className}
      stroke="currentColor"
      fill="currentColor"
      strokeWidth="0"
      viewBox="0 0 320 512"
      height="1em"
      width="1em"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M41 288h238c21.4 0 32.1 25.9 17 41L177 448c-9.4 9.4-24.6 9.4-33.9 0L24 329c-15.1-15.1-4.4-41 17-41zm255-105L177 64c-9.4-9.4-24.6-9.4-33.9 0L24 183c-15.1 15.1-4.4 41 17 41h238c21.4 0 32.1-25.9 17-41z"></path>
    </svg>
  );
}

export function SortUpIcon({ className }) {
  return (
    <svg
      className={className}
      stroke="currentColor"
      fill="currentColor"
      strokeWidth="0"
      viewBox="0 0 320 512"
      height="1em"
      width="1em"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M279 224H41c-21.4 0-32.1-25.9-17-41L143 64c9.4-9.4 24.6-9.4 33.9 0l119 119c15.2 15.1 4.5 41-16.9 41z"></path>
    </svg>
  );
}

export function SortDownIcon({ className }) {
  return (
    <svg
      className={className}
      stroke="currentColor"
      fill="currentColor"
      strokeWidth="0"
      viewBox="0 0 320 512"
      height="1em"
      width="1em"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M41 288h238c21.4 0 32.1 25.9 17 41L177 448c-9.4 9.4-24.6 9.4-33.9 0L24 329c-15.1-15.1-4.4-41 17-41z"></path>
    </svg>
  );
}

export function IssueIcon({ value, size }) {
  switch (value) {
    case "bug": {
      return (
        <AiFillBug className="bg-red-500 text-white rounded" size={size} />
      );
    }
    case "task": {
      return <BiTask className="bg-sky-500 text-white rounded" size={size} />;
    }
    case "subtask": {
      return <BsSubtract className="text-blue-500 rounded" size={size} />;
    }
    case "story": {
      return (
        <BsFillBookmarkFill
          className="bg-lime-400 text-white rounded"
          size={size}
        />
      );
    }
    case "epic": {
      return (
        <BsFillLightningFill
          className="bg-purple-500 text-white rounded"
          size={size}
        />
      );
    }
    default: {
      return <div>{value}</div>;
    }
  }
}
