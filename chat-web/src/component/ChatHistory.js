import React from 'react'
import Message from '../component/Message'

const ChatHistory = ({chatHistory}) => {

	console.log(chatHistory)

	const messages = chatHistory.map((msg,i) => <Message key={i} msg={msg} />);

	return (
		<div className='ChatHistory'>
			<h2>Chat History</h2>
			{messages}
		</div>
	);
}

export default ChatHistory
