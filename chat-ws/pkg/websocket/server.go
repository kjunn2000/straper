package websocket

import "fmt"

func (p *Pool) StartWSServer() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("New user enter room...")
			for client, _ := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Content: "New user enter room..."})
			}
			ch, err := p.Rdb.GetChatHistory()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Unable to push chat history")
				continue
			}
			msg := UpdateHistoryMessage{
				Type:    4,
				Content: ch,
			}
			client.Conn.WriteJSON(msg)

		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("One user quit room...")
			for client, _ := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 2, Content: "One user quit room..."})
			}

		case msg := <-p.Broadcast:
			fmt.Println("Receive message :" + msg.Content)
			err := p.Rdb.UpdateChatHistory(msg.Content)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for client, _ := range p.Clients {
				client.Conn.WriteJSON(msg)
			}
		}
	}
}
