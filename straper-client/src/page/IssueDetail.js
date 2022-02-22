import React, { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import SubPage from "../components/border/SubPage";
import { IssueIcon } from "../shared/Icons";
import useIssueStore from "../store/issueStore";
import { useHistory } from "react-router-dom";
import { convertToDateString } from "../service/object";
import EditIssueBtn from "../components/bug/EditIssueBtn";
import DeleteIssueBtn from "../components/bug/DeleteIssueBtn";

const IssueDetail = () => {
  const { issueId } = useParams();
  const issues = useIssueStore((state) => state.issues);
  const [issue, setIssue] = useState();
  const history = useHistory();

  useEffect(() => {
    getIssueData();
  }, [issueId]);

  const getIssueData = () => {
    setIssue(issues.find((i) => i.issue_id === issueId));
  };

  const detailHeaders = [
    "Type",
    "Backlog Priority",
    "Label",
    "Status",
    "Story Point",
    "Environment",
    "Serverity",
    "Reporter",
  ];
  const detailFields = [
    "type",
    "backlog_priority",
    "label",
    "status",
    "story_point",
    "environment",
    "serverity",
    "reporter",
  ];

  return (
    <SubPage>
      {issue && (
        <div className="flex flex-col space-y-5">
          <div className="flex p-2 space-x-3">
            <IssueIcon value={issue.type} size={50} />
            <div className="flex flex-col">
              <div
                className="text-sm text-blue-500 hover:cursor-pointer hover:font-semibold"
                onClick={() => history.push("/bug")}
              >
                HOME
              </div>
              <div className="text-2xl font-bold">{issue.summary}</div>
            </div>
          </div>
          <div className="flex space-x-5">
            <EditIssueBtn />
            <DeleteIssueBtn issueId={issue.issue_id} />
          </div>
          <div className="grid grid-cols-4 p-1">
            <div className="col-span-3 flex flex-col space-y-5 p-3">
              <div>
                <div className="font-bold">Details</div>
                <div className="flex space-x-5">
                  <div>
                    {detailHeaders.map((title) => (
                      <div className="font-semibold text-gray-500">
                        {title}:
                      </div>
                    ))}
                  </div>
                  <div>
                    {detailFields.map((field) => (
                      <div>{issue[field] || "-"}</div>
                    ))}
                  </div>
                </div>
              </div>
              <div className="group">
                <div>
                  <span className="font-bold">Description</span>
                </div>
                <div className="p-3 bg-gray-100 rounded break-all">
                  {issue.description || "-"}
                </div>
              </div>
              <div>
                <div className="font-bold">Attachments</div>
              </div>
              <div>
                <div className="font-bold">Acceptance Criteria</div>
                <div className="p-3 bg-gray-100 rounded break-all">
                  {issue.acceptance_criteria || "-"}
                </div>
              </div>
              <div>
                <div className="font-bold">Replicate Step</div>
                <div className="p-3 bg-gray-100 rounded break-all">
                  {issue.replicate_step || "-"}
                </div>
              </div>
              <div>
                <div className="font-bold">Workaround</div>
                <div className="p-3 bg-gray-100 rounded break-all">
                  {issue.workaround || "-"}
                </div>
              </div>
            </div>
            <div className="col-span-1 flex flex-col space-y-5">
              <div>
                <div className="font-bold">People</div>
                <div className="flex space-x-5">
                  <div className="font-semibold text-gray-500">
                    <div>Assignee:</div>
                  </div>
                  <div>{issue.assignee}</div>
                </div>
              </div>
              <div>
                <div className="font-bold">Dates</div>
                <div className="flex space-x-5">
                  <div className="font-semibold text-gray-500 flex flex-col">
                    <div>Created:</div>
                    <div>Due:</div>
                  </div>
                  <div className="flex flex-col">
                    <div>{convertToDateString(issue.created_date)}</div>
                    <div>{convertToDateString(issue.due_time)}</div>
                  </div>
                </div>
              </div>
              {issue.epic_link && issue.epic_link !== "" && (
                <div>
                  <div className="font-bold">Epic Link</div>
                  <Link
                    to={`/issue/${issue.epic_link}`}
                    className="text-blue-500 hover:text-blue-300 hover:cursor-pointer"
                  >
                    view
                  </Link>
                </div>
              )}
            </div>
          </div>
        </div>
      )}
    </SubPage>
  );
};

export default IssueDetail;
