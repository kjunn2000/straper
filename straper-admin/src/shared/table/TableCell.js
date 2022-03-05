import classNames from "classnames";
import React from "react";
import { convertToDateString } from "../../service/object";
import { TiTick } from "react-icons/ti";

export function DateCell({ value }) {
  return <div className="text-gray-500">{convertToDateString(value)}</div>;
}

export function StatusPill({ value }) {
  const status = value ? value.toUpperCase() : "UNKNOWN";

  return (
    <span
      className={classNames(
        "px-3 py-1 uppercase leading-wide font-bold text-xs rounded-full shadow-sm",
        status === "ACTIVE" ? "bg-green-100 text-green-800" : null,
        status === "VERIFYING" ? "bg-yellow-100 text-yellow-800" : null,
        status === "INACTIVE" ? "bg-red-100 text-red-800" : null
      )}
    >
      {status}
    </span>
  );
}

export function ActionCell({ column, row }) {
  return (
    <div className="inline-flex rounded-md shadow-sm" role="group">
      <button
        type="button"
        className="py-2 px-4 text-sm font-medium text-white bg-green-500 rounded-l-lg border border-gray-200 hover:bg-green-400"
        onClick={() => column.editAction(row.original[column.idAccessor])}
      >
        Edit
      </button>
      <button
        type="button"
        className="py-2 px-4 text-sm font-medium text-white bg-red-500 rounded-r-md border border-gray-200 hover:bg-red-400"
        onClick={() => column.deleteAction(row.original[column.idAccessor])}
      >
        Delete
      </button>
    </div>
  );
}

export function IsDefaultCell({ value }) {
  return value ? <TiTick className="text-green-500" size={20} /> : <></>;
}
