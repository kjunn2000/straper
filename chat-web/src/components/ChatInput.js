import React from 'react'
import useWorkspaceStore from "../store/workspaceStore"

const ChatInput = () => {
  const currChannel = useWorkspaceStore((state) => state.currChannel);

	return (
		<div className="bg-gray-800 bg-opacity-40 rounded-lg w-full">
        		<input className="bg-transparent p-3 w-full" placeholder={"Message #" + currChannel?.channel_name}/>
		</div>
	)
}

export default ChatInput
