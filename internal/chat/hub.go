package chat

import (
	"fmt"
)

var userNum = 0;

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}


func (h *Hub) Run(){
	for{
		select{
		case client :=<-h.register:
			h.clients[client] = true
			userNum +=1
			 fmt.Printf("User joined the chat! Users in chat: %d\n",userNum)

			case client := <- h.unregister:
				if _, ok := h.clients[client];ok{
					userNum-=1
					fmt.Printf("User left the chat! Users in chat: %d\n",userNum)
					delete(h.clients,client)
					close(client.send)
				}
			case message := <-h.broadcast:
				for client :=range h.clients{
					select{
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients,client)
					}
				}

		}
	}
}