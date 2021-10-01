import React from 'react'

const SidebarIcon = ({workspace,changeWorkspace}) => {

	return (
		<div className="w-auto text-center py-3">
			<button className="rounded-full text-white text-center h-12 w-12 items-center justify-center
			bg-gray-500 hover:bg-gray-800" onClick={()=>changeWorkspace(workspace.workspace_id)}>
				{workspace.workspace_name[0]}
			</button>
		</div>
	)
}

export default SidebarIcon
