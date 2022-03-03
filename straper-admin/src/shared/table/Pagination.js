import React, { useEffect, useState } from "react";
import { Button, PageButton } from "../Button";
import { HiChevronLeft, HiChevronRight } from "react-icons/hi";

const Pagination = ({
  pageChangeHandler,
  totalRows,
  rowsPerPage,
  refreshPage,
}) => {
  const noOfPages = Math.ceil(totalRows / rowsPerPage);

  const [currentPage, setCurrentPage] = useState(1);
  const [canGoBack, setCanGoBack] = useState(false);
  const [canGoNext, setCanGoNext] = useState(true);

  useEffect(() => {
    setCurrentPage(1);
  }, [refreshPage]);

  useEffect(() => {
    if (noOfPages === currentPage) {
      setCanGoNext(false);
    } else {
      setCanGoNext(true);
    }
    if (currentPage === 1) {
      setCanGoBack(false);
    } else {
      setCanGoBack(true);
    }
  }, [noOfPages, currentPage]);

  const handlePageChange = (targetPage, isNext) => {
    setCurrentPage(targetPage);
    pageChangeHandler(isNext);
  };

  const onNextPage = () => handlePageChange(currentPage + 1, true);
  const onPrevPage = () => handlePageChange(currentPage - 1, false);

  return (
    <>
      {noOfPages > 1 ? (
        <div className="p-5 flex items-center justify-between">
          <div className="flex-1 flex justify-between sm:hidden">
            <Button onClick={onPrevPage} disabled={!canGoBack}>
              Previous
            </Button>
            <Button onClick={onNextPage} disabled={!canGoNext}>
              Next
            </Button>
          </div>
          <div className="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div className="flex gap-x-2 items-baseline">
              <span className="text-sm text-gray-700">
                Page <span className="font-medium">{currentPage}</span> of{" "}
                <span className="font-medium">{noOfPages}</span>
              </span>
            </div>
            <div>
              <nav
                className="relative z-0 inline-flex rounded-md shadow-sm -space-x-px"
                aria-label="Pagination"
              >
                <PageButton onClick={onPrevPage} disabled={!canGoBack}>
                  <span className="sr-only">Previous</span>
                  <HiChevronLeft
                    className="h-5 w-5 text-gray-400"
                    aria-hidden="true"
                  />
                </PageButton>
                <PageButton onClick={onNextPage} disabled={!canGoNext}>
                  <span className="sr-only">Next</span>
                  <HiChevronRight
                    className="h-5 w-5 text-gray-400"
                    aria-hidden="true"
                  />
                </PageButton>
              </nav>
            </div>
          </div>
        </div>
      ) : null}
    </>
  );
};

export default Pagination;
