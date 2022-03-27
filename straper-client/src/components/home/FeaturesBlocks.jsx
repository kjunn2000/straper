import React from "react";
import {
  BsFillChatDotsFill,
  BsFillPeopleFill,
  BsFillHandbagFill,
} from "react-icons/bs";
import { BiTask } from "react-icons/bi";
import { AiFillBug, AiFillFile } from "react-icons/ai";

function FeaturesBlocks() {
  return (
    <section className="relative">
      {/* Section background (needs .relative class on parent and next sibling elements) */}
      <div
        className="absolute inset-0 top-1/2 md:mt-24 lg:mt-0 bg-slate-700 pointer-events-none"
        aria-hidden="true"
      ></div>
      <div className="absolute left-0 right-0 bottom-0 m-auto w-px p-px h-20 bg-gray-200 transform translate-y-1/2"></div>

      <div className="relative max-w-6xl mx-auto px-4 sm:px-6">
        <div className="py-12 md:py-20">
          {/* Section header */}
          <div className="max-w-3xl mx-auto text-center pb-12 md:pb-20">
            <h2 className="h2 mb-4">The Features of Straper</h2>
            <p className="text-xl text-gray-600">
              Any Thing You Need For Online Collaboration
            </p>
          </div>

          {/* Items */}
          <div className="max-w-sm mx-auto grid gap-6 md:grid-cols-2 lg:grid-cols-3 items-start md:max-w-2xl lg:max-w-none">
            {/* 1st item */}
            <div className="relative flex flex-col items-center p-6 bg-white rounded shadow-xl">
              <BsFillChatDotsFill className="w-16 h-16 p-1 -mt-1 mb-2 text-sky-600" />
              <h4 className="text-xl font-bold leading-snug tracking-tight mb-1">
                Real-time Chatting
              </h4>
              <p className="text-gray-600 text-center">
                Project members can communicate through channel.
              </p>
            </div>

            {/* 2nd item */}
            <div className="relative flex flex-col items-center p-6 bg-white rounded shadow-xl">
              <BiTask className="w-16 h-16 p-1 -mt-1 mb-2 text-slate-500" />
              <h4 className="text-xl font-bold leading-snug tracking-tight mb-1">
                Task Allocation
              </h4>
              <p className="text-gray-600 text-center">
                A drop & drop task board provides the best experience.
              </p>
            </div>

            {/* 3rd item */}
            <div className="relative flex flex-col items-center p-6 bg-white rounded shadow-xl">
              <AiFillBug className="w-16 h-16 p-1 -mt-1 mb-2 text-red-600" />
              <h4 className="text-xl font-bold leading-snug tracking-tight mb-1">
                Bug Tracking
              </h4>
              <p className="text-gray-600 text-center">
                Manage project bugs in a straigtforward and easy way.
              </p>
            </div>

            {/* 4th item */}
            <div className="relative flex flex-col items-center p-6 bg-white rounded shadow-xl">
              <BsFillHandbagFill className="w-16 h-16 p-1 -mt-1 mb-2 text-slate-600" />
              <h4 className="text-xl font-bold leading-snug tracking-tight mb-1">
                Workspace Management
              </h4>
              <p className="text-gray-600 text-center">
                Manage multiple projects with the concept of online workspace.
              </p>
            </div>

            {/* 5th item */}
            <div className="relative flex flex-col items-center p-6 bg-white rounded shadow-xl">
              <BsFillPeopleFill className="w-16 h-16 p-1 -mt-1 mb-2 text-indigo-600" />
              <h4 className="text-xl font-bold leading-snug tracking-tight mb-1">
                User Status
              </h4>
              <p className="text-gray-600 text-center">
                Track the project team member's status on the fly.
              </p>
            </div>

            {/* 6th item */}
            <div className="relative flex flex-col items-center p-6 bg-white rounded shadow-xl">
              <AiFillFile className="w-16 h-16 p-1 -mt-1 mb-2 text-gray-500" />
              <h4 className="text-xl font-bold leading-snug tracking-tight mb-1">
                File Sharing
              </h4>
              <p className="text-gray-600 text-center">
                Upload the attachments to Straper and share with other project
                members.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default FeaturesBlocks;
