import React from 'react'

const ChatInput = ({channel}) => {
	return (
		<div className="bg-gray-800 bg-opacity-40 rounded-lg w-full">
        		<input className="bg-transparent p-3 w-full" placeholder={"Message #" + channel.channel_name}/>
		</div>
	)
}

export default ChatInput
