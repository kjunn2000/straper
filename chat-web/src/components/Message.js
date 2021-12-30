import React from 'react'

const Message = ({msg}) => {
	
	return (
		<div className="text-white flex flex-row space-x-6 p-4">
			<div className="rounded-full text-white text-center h-10 w-10 flex items-center justify-center 
			bg-gray-500">
				{msg.creator_name[0]}
			</div>
			<div className="flex flex-col">
				<div className="text-lg font-medium">{msg.creator_name}</div>
				<div className="text-gray-300">{msg.content}</div>
			</div>
		</div>
	);

}

export default Message
