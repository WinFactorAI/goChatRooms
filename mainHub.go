package main

import (
	"sync"
)

// 主Hub管理所有房间
type MainHub struct {
	Rooms map[string]*Room // 所有房间映射表
	Mutex sync.RWMutex     // 房间操作锁
}

func NewMainHub() *MainHub {
	return &MainHub{
		Rooms: make(map[string]*Room),
	}
}

func (h *MainHub) CreateRoom(roomID string) *Room {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	if _, ok := h.Rooms[roomID]; !ok {
		room := &Room{
			ID:        roomID,
			Clients:   make(map[*Client]bool),
			Broadcast: make(chan []byte, 256),
		}
		h.Rooms[roomID] = room
		go room.run() // 启动房间消息处理
	}
	return h.Rooms[roomID]
}
func (h *MainHub) GarbageCollect() {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	for id, room := range h.Rooms {
		if len(room.Clients) == 0 {
			close(room.Broadcast)
			delete(h.Rooms, id)
		}
	}
}
