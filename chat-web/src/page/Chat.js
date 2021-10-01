import { useEffect, useState } from "react";
import ChatRoom from "../components/ChatRoom";
import Sidebar from "../components/Sidebar";

function Chat() {
  return (
    <div className="flex flex-row">
      <Sidebar />
      <ChatRoom />
    </div>
  );
}

export default Chat;
