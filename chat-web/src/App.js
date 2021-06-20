import './App.scss';
import Header from './component/Header';
import ChatHistory from './component/ChatHistory';
import ChatInput from './component/ChatInput';
import { useEffect, useState } from 'react';
import { connect, sendMsg } from "./api";


function App() {

  const [chatHistory, setChatHistory] = useState([])

  useEffect(()=> {
    connect(updateHistory)
  },[])


  const updateHistory = (msg) => {
    if (msg.type === 4){
      console.log(msg)
      setChatHistory(old => [...old, ...msg.content])
      return
    }
    setChatHistory(old => [...old,msg.content])
  }

  const send = (event) => {
    if (event.keyCode === 13 ) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <ChatInput send ={send}/>
    </div>
  );
}

export default App;
