package domain

import "fmt"

func (c *Channel) StartWSServer() {
	for {
		select {
		case client := <-c.Register:
			c.Clients[client] = true
			fmt.Println("New user enter room...")
			for client, _ := range c.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Content: "New user enter room..."})
			}
			ch, err := c.Rdb.GetChatHistory()
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

		case client := <-c.Unregister:
			delete(c.Clients, client)
			fmt.Println("One user quit room...")
			for client, _ := range c.Clients {
				client.Conn.WriteJSON(Message{Type: 2, Content: "One user quit room..."})
			}

		case msg := <-c.Broadcast:
			fmt.Println("Receive message :" + msg.Content)
			err := c.Rdb.UpdateChatHistory(msg.Content)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for client, _ := range c.Clients {
				client.Conn.WriteJSON(msg)
			}
		}
	}
}
