import React from 'react'
import WorkspaceMenu from './WorkspaceMenu'

function ChannelSidebar({workspace}){
	
	return (
		<div className="flex flex-col w-64 h-screen bg-gray-800">
    			<WorkspaceMenu workspace={workspace}/>
			<div className="px-3">
				{
					workspace?.channel_list && workspace.channel_list.map(channel=> (
						<div className="text-white font-medium p-3 hover:bg-gray-700 rounded-lg"
							key={channel?.channel_id}> # {channel?.channel_name}</div>
					))
				}
			</div>
  		</div>
	)
}

export default ChannelSidebar

