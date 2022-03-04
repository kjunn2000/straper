import classNames from "classnames";
import React, { useState } from "react";
import {
  useTable,
  useFilters,
  useGlobalFilter,
  useAsyncDebounce,
  useSortBy,
} from "react-table";
import { convertToDateString } from "../../service/object";
import { SortIcon, SortUpIcon, SortDownIcon } from "../../shared/Icons";
import Loader from "../Loader";

// Define a default UI for filtering

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

export function ActionCell({ value, column, row }) {
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

function Table({ columns, data, isLoading, totalCount, onSearch }) {
  const { getTableProps, getTableBodyProps, headerGroups, prepareRow, rows } =
    useTable(
      {
        columns,
        data,
        manualPaginations: true,
      },
      useSortBy
    );

  const [searchStr, setSearchStr] = useState("");

  const handleSearch = (value) => {
    setSearchStr(value);
    onSearch(value);
  };

  return (
    <>
      {isLoading ? (
        <Loader />
      ) : (
        <>
          <div className="sm:flex sm:gap-x-2 p-5">
            <label className="flex gap-x-2 items-baseline">
              <span className="text-gray-700 font-semibold">Search: </span>
              <input
                type="text"
                className="rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 p-1"
                placeholder={`${totalCount} records...`}
                onKeyDown={(e) =>
                  e.key === "Enter" && handleSearch(e.target.value)
                }
              />
            </label>
            {headerGroups.map((headerGroup) =>
              headerGroup.headers.map((column) =>
                column.Filter ? (
                  <div className="mt-2 sm:mt-0" key={column.id}>
                    {column.render("Filter")}
                  </div>
                ) : null
              )
            )}
          </div>
          {searchStr && searchStr !== "" && (
            <div className="text-gray-500 text-sm px-5">
              There are {totalCount} records for seach string "
              <span className="font-semibold italic text-indigo-600">
                {searchStr}
              </span>
              ".
            </div>
          )}
          <div className="mt-4 flex flex-col">
            <div className="-my-2 overflow-x-auto">
              <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
                <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
                  <table
                    {...getTableProps()}
                    className="min-w-full divide-y divide-gray-200"
                  >
                    <thead className="bg-gray-50">
                      {headerGroups.map((headerGroup) => (
                        <tr {...headerGroup.getHeaderGroupProps()}>
                          {headerGroup.headers.map((column) => (
                            // Add the sorting props to control sorting. For this example
                            // we can add them into the header props
                            <th
                              scope="col"
                              className="group px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                              {...column.getHeaderProps(
                                column.getSortByToggleProps()
                              )}
                            >
                              <div className="flex items-center justify-between">
                                {column.render("Header")}
                                {/* Add a sort direction indicator */}
                                <span>
                                  {column.isSorted ? (
                                    column.isSortedDesc ? (
                                      <SortDownIcon className="w-4 h-4 text-gray-400" />
                                    ) : (
                                      <SortUpIcon className="w-4 h-4 text-gray-400" />
                                    )
                                  ) : (
                                    <SortIcon className="w-4 h-4 text-gray-400 opacity-0 group-hover:opacity-100" />
                                  )}
                                </span>
                              </div>
                            </th>
                          ))}
                        </tr>
                      ))}
                    </thead>
                    <tbody
                      {...getTableBodyProps()}
                      className="bg-white divide-y divide-gray-200"
                    >
                      {rows.map((row, i) => {
                        // new
                        prepareRow(row);
                        return (
                          <tr {...row.getRowProps()}>
                            {row.cells.map((cell) => {
                              return (
                                <td
                                  {...cell.getCellProps()}
                                  className="px-6 py-4 whitespace-nowrap"
                                  role="cell"
                                >
                                  {cell.column.Cell.name ===
                                  "defaultRenderer" ? (
                                    <div className="text-sm text-gray-500">
                                      {cell.render("Cell")}
                                    </div>
                                  ) : (
                                    cell.render("Cell")
                                  )}
                                </td>
                              );
                            })}
                          </tr>
                        );
                      })}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </>
      )}
    </>
  );
}

export default Table;
