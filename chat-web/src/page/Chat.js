import { useEffect, useState } from 'react';
import ChatRoom from '../component/ChatRoom';
import Sidebar from '../component/Sidebar';


function Chat() {

  const [workspaces,setWorkspaces] = useState([
		{"workspace_id":"W000001",
		"workspace_name":"FreeStyle", 
		"channel_list":[{"channel_id":"C00001","channel_name":"General",},
		{"channel_id":"C00002","channel_name":"Leetcode"}]},

		{"workspace_id":"W000002",
		"workspace_name":"Google", 
		"channel_list":[{"channel_id":"C00003","channel_name":"General"},
		{"channel_id":"C00003","channel_name":"Meeting"}]},
	]);

  // const [chatHistory, setChatHistory] = useState([])

  // useEffect(()=> {
  //   connect(updateHistory)
  // },[])


  // const updateHistory = (msg) => {
  //   if (msg.type === 4){
  //     console.log(msg)
  //     setChatHistory(old => [...old, ...msg.content])
  //     return
  //   }
  //   setChatHistory(old => [...old,msg.content])
  // }

  // const send = (event) => {
  //   if (event.keyCode === 13 ) {
  //     sendMsg(event.target.value);
  //     event.target.value = "";
  //   }
  // }

  return (
    <div className="flex flex-row">
      <Sidebar/>
      <ChatRoom channel={workspaces[0].channel_list[0]}/>
    </div>
  );
}

export default Chat;

