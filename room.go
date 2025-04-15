package main

import "log"

// 房间结构体
type Room struct {
	ID        string           // 房间唯一标识
	Clients   map[*Client]bool // 房间内客户端
	Broadcast chan []byte      // 房间级消息通道
}

func (r *Room) run() {
	defer func() {
		close(r.Broadcast)
		log.Printf("房间 %s 已关闭", r.ID)
	}()

	for {

		select {
		case message := <-r.Broadcast:
			// 广播消息给房间内所有客户端

			for client := range r.Clients {
				select {
				case client.send <- message:
					log.Println("broadcast msg ", len(message))
				default:
					close(client.send)
					delete(r.Clients, client)
				}
			}

		}
	}
}
