import React from 'react'

const ChatInput = () => {
	return (
		<div className="bg-gray-800 bg-opacity-40 rounded-lg w-full">
        		<input className="bg-transparent p-3 w-full" placeholder={"Message #" + 
				// channel.channel_name
				"Hello"
				}/>
		</div>
	)
}

export default ChatInput
