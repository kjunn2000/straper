import React from 'react'

const ChatInput = ({send}) => {
	return (
		<div>
			<div className="ChatInput">
        			<input onKeyDown={send} />
      			</div>	
		</div>
	)
}

export default ChatInput
