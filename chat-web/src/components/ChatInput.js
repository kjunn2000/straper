import React from 'react'
import { sendMsg } from '../service/websocket';
import useIdentifyStore from '../store/identityStore';
import useWorkspaceStore from "../store/workspaceStore"

const ChatInput = () => {
  const currChannel = useWorkspaceStore((state) => state.currChannel);
  const identity = useIdentifyStore((state) => state.identity);

  const defaultPlaceHolder = "Message #" + currChannel?.channel_name;

  const handleKeyDown = (event) => {
	  if (event.key !== "Enter"){
		  return;
	  }
	  sendMsg(currChannel.channel_id, identity.username, event.target.value);
	  event.target.value = "";
  }

	return (
		<div className="bg-gray-800 bg-opacity-40 rounded-lg w-full">
        		<input className="bg-transparent p-3 w-full" 
				placeholder={defaultPlaceHolder}
				onKeyDown={(e)=>handleKeyDown(e)}
				/>
		</div>
	)
}

export default ChatInput
