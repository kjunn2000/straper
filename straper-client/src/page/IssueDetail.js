import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import SubPage from "../components/border/SubPage";
import { IssueIcon } from "../shared/Icons";
import useIssueStore from "../store/issueStore";
import { useHistory } from "react-router-dom";
import { convertToDateString } from "../service/object";

const IssueDetail = () => {
  const { issueId } = useParams();
  const issues = useIssueStore((state) => state.issues);
  const [issue, setIssue] = useState();
  const history = useHistory();

  useEffect(() => {
    getIssueData();
  }, []);

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
          <div className="grid grid-cols-4">
            <div className="col-span-3 flex flex-col space-y-5">
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
              <div>
                <div className="font-bold">Description</div>
                <div>{issue.description || "-"}</div>
              </div>
              <div>
                <div className="font-bold">Attachments</div>
              </div>
              <div>
                <div className="font-bold">Acceptance Criteria</div>
                <div>{issue.acceptance_criteria || "-"}</div>
              </div>
              <div>
                <div className="font-bold">Replicate Step</div>
                <div>{issue.replicate_step || "-"}</div>
              </div>
              <div>
                <div className="font-bold">Workaround</div>
                <div>{issue.workaround || "-"}</div>
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
                  <div className="font-semibold text-gray-500">
                    <div>Created:</div>
                    <div>Due:</div>
                  </div>
                  <div>
                    {convertToDateString(issue.created_date)}
                    {convertToDateString(issue.due_time)}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </SubPage>
  );
};

export default IssueDetail;
